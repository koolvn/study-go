package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/koolvn/study-go.git/auth"
)

type VerifyRequest struct {
	Token string `json:"token"`
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	logRequest("[INFO] Received request for /", r)
	handler := http.FileServer(http.Dir("static"))
	handler.ServeHTTP(w, r)
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	logRequest("[INFO] Received auth request", r)

	var authRequest auth.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		msg := "[ERROR] " + err.Error()
		log.Println(msg)
		writeJSONResponse(http.StatusBadRequest, msg, w)
		return
	}
	defer r.Body.Close()

	if authRequest.Username == "" || authRequest.Password == "" {
		writeJSONResponse(http.StatusBadRequest, "username and password are required", w)
		return
	}

	isLdapAuthorized, errLDAP := auth.LdapAuthenticateUser(authRequest)
	if errLDAP != nil {
		msg := "[ERROR] " + errLDAP.Error()
		log.Println(msg)
		writeJSONResponse(http.StatusUnauthorized, msg, w)
		return
	}
	if !isLdapAuthorized {
		writeJSONResponse(http.StatusUnauthorized, "LDAP authorization failed", w)
		return
	}
	token, err := auth.CreateToken(authRequest.Username)
	if err != nil {
		msg := "[ERROR] " + err.Error()
		log.Println(msg)
		writeJSONResponse(http.StatusInternalServerError, msg, w)
		return
	}

	writeJSONResponse(http.StatusOK, map[string]string{"token": token}, w)
}

func HandleAuthVerify(w http.ResponseWriter, r *http.Request) {
	logRequest("[INFO] Received auth verify request", r)

	var verifyRequest VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&verifyRequest); err != nil {
		msg := "[ERROR] " + err.Error()
		log.Println(msg)
		writeJSONResponse(http.StatusBadRequest, msg, w)
		return
	}
	defer r.Body.Close()

	if verifyRequest.Token == "" {
		writeJSONResponse(http.StatusUnauthorized, "No token provided", w)
		return
	}

	if err := auth.VerifyToken(verifyRequest.Token); err != nil {
		msg := "[ERROR] " + err.Error()
		log.Println(msg)
		writeJSONResponse(http.StatusUnauthorized, msg, w)
	} else {
		writeJSONResponse(http.StatusOK, "Token verified!", w)
	}
}
