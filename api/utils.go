package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONResponse struct {
	StatusCode int `json:"status"`
	Details    any `json:"details"`
}

func writeJSONResponse(statusCode int, details any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := JSONResponse{StatusCode: statusCode, Details: details}

	jsonResp, _ := json.Marshal(resp)
	_, err := w.Write(jsonResp)
	if err != nil {
		log.Printf("[ERROR] Failed to write JSON response: %s", err)
	}
}
