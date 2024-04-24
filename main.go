package main

import (
	"github.com/koolvn/study-go.git/auth"
	log "golang.org/x/exp/slog"
	"os"
)

func main() {
	// mock user data
	user := auth.UserLogin{
		Username: "exampleUser",
		Password: "examplePassword",
	}

	conn, err := auth.Connect()
	if err != nil {
		log.Error("LDAP connection failed.")
		os.Exit(1)
	}

	defer conn.Close()

	authenticated, authErr := auth.Auth(conn, user)
	if authErr != nil {
		log.Error("Authentication failed.")
		os.Exit(1)
	}

	if authenticated {
		log.Info("User authenticated successfully.")
	} else {
		log.Warn("Authentication failed. Invalid credentials.")
	}
}
