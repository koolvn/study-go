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
		msg := fmt.Sprintf("LDAP connection failed, error details: %v", err)
		log.Error(msg)
		return nil, err
	}

	if err := conn.Bind(os.Getenv("BIND_USER"), os.Getenv("BIND_PASSWORD")); err != nil {
		msg := fmt.Sprintf("LDAP bind failed while connecting, error details: %v", err)
		log.Error(msg)
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
		msg := fmt.Sprintf("LDAP search failed for user `%s`, error details: %v", user.Username, err)
		log.Error(msg)
		return false, err
	}

	if len(searchResp.Entries) != 1 {
		msg := fmt.Sprintf("User `%s` not found or multiple entries found", user.Username)
		log.Error(msg)
		err = fmt.Errorf(msg)
		return false, err
	}
	msg := fmt.Sprintf("User `%s` found", user.Username)
	log.Info(msg)
	userDN := searchResp.Entries[0].DN
	err = conn.Bind(userDN, user.Password)
	if err != nil {
		msg := fmt.Sprintf("LDAP authentication failed for user `%s`, error details: %v", user.Username, err)
		log.Error(msg)
		err = fmt.Errorf(msg)
		return false, err
	}
	return true, nil
}

func AuthenticateUser(username string, password string) (bool, error) {
	user := UserLogin{
		Username: username,
		Password: password,
	}

	conn, err := Connect()
	if err != nil {
		return false, err
	}

	defer conn.Close()

	authenticated, authErr := Auth(conn, user)
	if authErr != nil {
		return false, authErr
	}

	if authenticated {
		log.Info("User authenticated successfully.")
	} else {
		log.Warn("Authentication failed. Invalid credentials.")
	}
	return authenticated, nil
}
