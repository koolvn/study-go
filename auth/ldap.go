package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/go-ldap/ldap"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LdapAuthenticateUser authenticates a user against an LDAP server.
// It takes a UserLogin struct as input, which contains the user's credentials,
// and returns a boolean indicating whether the authentication was successful,
// along with an error if the authentication process fails at any point.
// The function establishes a connection to the LDAP server using LdapConnect,
// and defers the closing of this connection until the function completes.
// It then attempts to authenticate the user with the provided credentials
// using the LdapAuth function. If any step fails, the function returns false
// and the error encountered. Otherwise, it returns true and nil error upon
// successful authentication.
func LdapAuthenticateUser(user UserLogin) (bool, error) {
	conn, err := LdapConnect()
	if err != nil {
		return false, err
	}

	defer conn.Close()

	authenticated, authErr := LdapAuth(conn, user)
	if authErr != nil {
		return false, authErr
	}
	return authenticated, nil
}

// LdapConnect establishes a connection to the LDAP server.
func LdapConnect() (*ldap.Conn, error) {
	conn, err := ldap.DialURL(os.Getenv("LDAP_ADDRESS"))
	if err != nil {
		msg := fmt.Sprintf("[ERROR] LDAP connection failed, error details: %v", err)
		log.Println(msg)
		return nil, err
	}

	if err := conn.Bind(os.Getenv("BIND_USER"), os.Getenv("BIND_PASSWORD")); err != nil {
		msg := fmt.Sprintf("[ERROR] LDAP bind failed while connecting, error details: %v", err)
		log.Println(msg)
		return nil, err
	}

	return conn, nil
}

// LdapAuth performs LDAP authentication for the provided user credentials.
func LdapAuth(conn *ldap.Conn, user UserLogin) (bool, error) {
	searchRequest := ldap.NewSearchRequest(
		os.Getenv("LDAP_BASE_DN"),
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(sAMAccountName=%s)", user.Username),
		[]string{"dn"},
		nil,
	)

	searchResp, err := conn.Search(searchRequest)
	if err != nil {
		msg := fmt.Sprintf(
			"[ERROR] LDAP search failed for user `%s`, error details: %v", user.Username, err)
		log.Println(msg)
		return false, err
	}

	if len(searchResp.Entries) != 1 {
		msg := fmt.Sprintf(
			"[ERROR] User `%s` not found or multiple entries found", user.Username)
		log.Println(msg)
		err = fmt.Errorf(msg)
		return false, err
	}
	msg := fmt.Sprintf("[INFO] User `%s` found", user.Username)
	log.Println(msg)
	userDN := searchResp.Entries[0].DN
	err = conn.Bind(userDN, user.Password)
	if err != nil {
		msg := fmt.Sprintf(
			"[ERROR] LDAP authentication failed for user `%s`, error details: %v",
			user.Username, err)
		log.Println(msg)
		err = fmt.Errorf(msg)
		return false, err
	}
	return true, nil
}
