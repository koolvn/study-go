package api

import (
	"encoding/json"
	"fmt"
	"github.com/koolvn/study-go.git/auth"
	"github.com/koolvn/study-go.git/types"
	"github.com/koolvn/study-go.git/utils"
	log "golang.org/x/exp/slog"
	"net/http"
)

func AuthPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug(fmt.Sprintf("Received auth page request from %v", r.RemoteAddr))
	handler := http.FileServer(http.Dir("static"))
	handler.ServeHTTP(w, r)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var creds types.Credentials
	w.Header().Set("Content-Type", "application/json")

	log.Debug(fmt.Sprintf("Received auth request from %v", r.RemoteAddr))
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.PrepareResponse("Unauthorized", err.Error(), w)
		return
	}

	authorized, err := auth.AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.PrepareResponse("Unauthorized", err.Error(), w)
		return
	}

	if authorized {
		log.Info(fmt.Sprintf("User %v authenticated: %v", creds.Username, authorized))
		utils.PrepareResponse("Authorized", "", w)
	} else {
		msg := fmt.Sprintf("User %v authentication failed", creds.Username)
		log.Info(msg)
		w.WriteHeader(http.StatusUnauthorized)
		utils.PrepareResponse("Unauthorized", msg, w)
	}
	w.WriteHeader(http.StatusOK)
}
