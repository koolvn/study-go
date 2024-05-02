package api

import (
	"encoding/json"
	"net/http"

	"github.com/koolvn/study-go.git/auth"
)

type VerifyRequest struct {
	Token string `json:"token"`
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	logRequest("[INFO] Received request for /", r)
	writeJSONResponse(http.StatusOK, "Hello World", w)
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	logRequest("[INFO] Received auth request", r)

	var authRequest auth.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		writeError(w, "Bad request", http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if authRequest.Username == "" || authRequest.Password == "" {
		writeError(w, "Bad request: username and password are required", http.StatusBadRequest, nil)
		return
	}

	isLdapAuthorized, errLDAP := auth.LdapAuthenticateUser(authRequest)
	if errLDAP != nil {
		writeError(w, errLDAP.Error(), http.StatusBadRequest, nil)
		return
	}
	if !isLdapAuthorized {
		writeError(w, "LDAP authorization failed", http.StatusUnauthorized, nil)
		return
	}
	token, err := auth.CreateToken(authRequest.Username)
	if err != nil {
		writeError(w, "Internal server error", http.StatusInternalServerError, err)
		return
	}

	writeJSONResponse(http.StatusOK, map[string]string{"token": token}, w)
}

func HandleAuthVerify(w http.ResponseWriter, r *http.Request) {
	logRequest("[INFO] Received auth verify request", r)

	var verifyRequest VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&verifyRequest); err != nil {
		writeError(w, "Bad request", http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if verifyRequest.Token == "" {
		writeError(w, "No token provided", http.StatusUnauthorized, nil)
		return
	}

	if err := auth.VerifyToken(verifyRequest.Token); err != nil {
		writeError(w, "Unauthorized", http.StatusUnauthorized, err)
	} else {
		writeJSONResponse(http.StatusOK, "Token verified!", w)
	}
}
