package auth

import (
	"fmt"

	"github.com/go-ldap/ldap"
	"github.com/koolvn/study-go.git/config"
)

type LDAPUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LDAPAuthenticator struct {
	cfg config.Config
}

func NewLDAPAuthenticator() *LDAPAuthenticator {
	return &LDAPAuthenticator{cfg: *config.NewConfig()}
}

// AuthorizeUser attempts to authenticate a given LDAPUser against the LDAP server.
func (a LDAPAuthenticator) AuthorizeUser(user LDAPUser) (bool, error) {
	conn, err := a.connect()
	if err != nil {
		return false, err
	}

	defer conn.Close()

	authenticated, authErr := a.authorize(conn, user)
	if authErr != nil {
		return false, authErr
	}
	return authenticated, nil
}

// connect establishes a connection to the LDAP server.
func (a LDAPAuthenticator) connect() (*ldap.Conn, error) {
	conn, errConnect := ldap.DialURL(a.cfg.LdapAddr)
	if errConnect != nil {
		return nil, errConnect
	}

	if err := conn.Bind(a.cfg.LdapBindUser, a.cfg.LdapBindPass); err != nil {
		return nil, err
	}

	return conn, nil
}

// authorize performs LDAP authentication for the provided user credentials.
func (a LDAPAuthenticator) authorize(conn *ldap.Conn, user LDAPUser) (bool, error) {
	searchRequest := ldap.NewSearchRequest(
		a.cfg.LdapBaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(sAMAccountName=%s)", user.Username),
		[]string{"dn"},
		nil,
	)

	searchResp, err := conn.Search(searchRequest)
	if err != nil {
		return false, err
	}

	if len(searchResp.Entries) != 1 {
		msg := fmt.Sprintf(
			"user `%s` not found or multiple entries found", user.Username)
		err = fmt.Errorf(msg)
		return false, err
	}
	userDN := searchResp.Entries[0].DN
	err = conn.Bind(userDN, user.Password)
	if err != nil {
		msg := fmt.Sprintf(
			"LDAP authentication failed for user `%s`, error details: %v",
			user.Username, err)
		err = fmt.Errorf(msg)
		return false, err
	}
	return true, nil
}
