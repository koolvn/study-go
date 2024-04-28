package api

import (
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
