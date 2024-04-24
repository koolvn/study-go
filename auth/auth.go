package auth

import (
	"fmt"
	"os"

	"github.com/go-ldap/ldap"
	log "golang.org/x/exp/slog"
)

// UserLogin struct represents the user credentials.
type UserLogin struct {
	Username string
	Password string
}

// Connect establishes a connection to the LDAP server.
func Connect() (*ldap.Conn, error) {
	conn, err := ldap.DialURL(os.Getenv("LDAP_ADDRESS"))
	if err != nil {
		log.Error("LDAP connection failed, error details: %v", err)
		return nil, err
	}

	if err := conn.Bind(os.Getenv("BIND_USER"), os.Getenv("BIND_PASSWORD")); err != nil {
		log.Error("LDAP bind failed while connecting, error details: %v", err)
		return nil, err
	}

	return conn, nil
}

// Auth performs LDAP authentication for the provided user credentials.
func Auth(conn *ldap.Conn, user UserLogin) (bool, error) {
	searchRequest := ldap.NewSearchRequest(
		os.Getenv("LDAP_BASE_DN"),
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(sAMAccountName=%s)", user.Username),
		[]string{"dn"},
		nil,
	)

	searchResp, err := conn.Search(searchRequest)
	if err != nil {
		log.Error("LDAP search failed for user %s, error details: %v", user.Username, err)
		return false, err
	}

	if len(searchResp.Entries) != 1 {
		log.Error("User: %s not found or multiple entries found", user.Username)
		err = fmt.Errorf("user: %s not found or multiple entries found", user.Username)
		return false, err
	}
	log.Info("User: %s found", user.Username)
	userDN := searchResp.Entries[0].DN
	err = conn.Bind(userDN, user.Password)
	if err != nil {
		log.Error("LDAP authentication failed for user %s, error details: %v", user.Username, err)
		err = fmt.Errorf("LDAP authentication failed for user %s", user.Username)
		return false, err
	}

	return true, nil
}
