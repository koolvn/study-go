package main

import (
	"flag"
	"github.com/koolvn/study-go.git/structs"
	"github.com/koolvn/study-go.git/utils"
	"html/template"
	"log"
	"net/http"
)

func main() {
	var host = flag.String("host", "0.0.0.0", "host to listen on")
	var port = flag.String("port", "8080", "port to listen on")
	flag.Parse()

	addr := *host + ":" + *port
	log.Printf("Listening on http://%s\n", addr)
	mux := createMux()
	log.Fatal(http.ListenAndServe(addr, mux))
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	mux.HandleFunc("GET /", serveRoot)
	return mux
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
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
