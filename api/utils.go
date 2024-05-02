package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JSONResponse struct {
	StatusCode int `json:"status"`
	Details    any `json:"details"`
}

// writeJSONResponse sends a JSON response with the specified status code and details.
func writeJSONResponse(statusCode int, details any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	resp := JSONResponse{StatusCode: statusCode, Details: details}
	jsonResp, errMarshal := json.Marshal(resp)
	if errMarshal != nil {
		log.Printf("[ERROR] Failed to marshal JSON response: %s", errMarshal)
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("[ERROR] Failed to write JSON response: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func logRequest(message string, r *http.Request) {
	log.Printf("%s from %s", message, r.RemoteAddr)
}

func writeError(w http.ResponseWriter, message string, statusCode int, err error) {
	if err != nil {
		log.Printf("[ERROR] %s", err)
		http.Error(w, fmt.Sprintf("%s: %s", message, err), statusCode)
	} else {
		http.Error(w, message, statusCode)
	}
}
