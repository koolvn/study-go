package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/koolvn/study-go.git/structs"
	"github.com/koolvn/study-go.git/utils"
)

func ServeRoot(w http.ResponseWriter, r *http.Request) {
	browser := utils.GetBrowserName(r.Header.Get("User-Agent"))
	user := structs.User{
		Ip:      r.RemoteAddr,
		Browser: browser,
	}
	log.Printf("[INFO]  GET / from %s", user)
	tmpl, err := template.ParseFiles("html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReceiveImage(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO]  POST /image from %s  Content-Length: %s", r.RemoteAddr, r.Header.Get("Content-Length"))
	w.WriteHeader(200)
}
