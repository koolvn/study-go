package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func PrepareResponse(status string, err string, w http.ResponseWriter) {
	response := map[string]string{}
	response["status"] = status
	response["error"] = err
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Printf("Failed to convert response to JSON: %s", jsonError)
	}
	_, writeErr := w.Write(jsonResponse)
	if writeErr != nil {
		log.Printf("Failed to write JSON response: %s", writeErr)
	}
}
