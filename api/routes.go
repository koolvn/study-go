package api

import (
	"encoding/json"
	"fmt"
	"github.com/koolvn/study-go.git/auth"
	"log"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
}

type VerifyRequest struct {
	Token string `json:"token"`
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] Received request for / from %s", r.RemoteAddr)
	writeJSONResponse(http.StatusOK, "Hello World", w)
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] Received auth request from %s", r.RemoteAddr)

	var authRequest AuthRequest
	errorJsonDecode := json.NewDecoder(r.Body).Decode(&authRequest)
	if errorJsonDecode != nil {
		message := fmt.Sprintf("Bad request. %s", errorJsonDecode.Error())
		writeJSONResponse(http.StatusBadRequest, message, w)
	}
	token, err := auth.CreateToken(authRequest.Username)
	if err != nil {
		writeJSONResponse(http.StatusInternalServerError, err.Error(), w)
	}

	msg := map[string]string{"token": token}
	writeJSONResponse(http.StatusOK, msg, w)

}

func HandleAuthVerify(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] Received auth verify request from %s", r.RemoteAddr)
	var verifyRequest VerifyRequest
	errorJsonDecode := json.NewDecoder(r.Body).Decode(&verifyRequest)
	if errorJsonDecode != nil {
		message := fmt.Sprintf("Bad request. %s", errorJsonDecode.Error())
		writeJSONResponse(http.StatusBadRequest, message, w)
	}
	token := verifyRequest.Token
	log.Println("TOKEN: ", token)
	if token == "" {
		writeJSONResponse(http.StatusUnauthorized, "No token provided", w)
	} else {
		err := auth.VerifyToken(token)
		if err != nil {
			writeJSONResponse(http.StatusUnauthorized, err.Error(), w)
		} else {
			writeJSONResponse(http.StatusOK, "Token verified!", w)
		}
	}
}
