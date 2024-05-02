package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type API struct {
	host     string
	port     string
	https    bool
	certPath string
	keyPath  string
	server   *http.Server
}

func NewAPI(host string, port string, https bool, certPath string, keyPath string) *API {
	return &API{
		host:     host,
		port:     port,
		https:    https,
		certPath: certPath,
		keyPath:  keyPath,
	}
}

func (api *API) Start() error {
	router := api.configureRoutes()
	serveAddr := fmt.Sprintf("%s:%s", api.host, api.port)

	api.server = &http.Server{
		Addr:    serveAddr,
		Handler: router,
	}

	// Handling graceful shutdown
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

		<-stop

		log.Println("[INFO] Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := api.server.Shutdown(ctx); err != nil {
			log.Fatalf("[ERROR] Error shutting down server: %v", err)
		} else {
			log.Println("[INFO] Server shutdown completed.")
		}

		log.Println("[INFO] Server gracefully stopped.")
	}()

	if api.https {
		log.Printf("[INFO] Starting server at https://%s\n", serveAddr)
		return api.server.ListenAndServeTLS(api.certPath, api.keyPath)
	} else {
		log.Printf("[INFO] Starting server at http://%s\n", serveAddr)
		log.Println("[WARNING] HTTPS is not enabled. It is recommended to use HTTPS in production.")
		return api.server.ListenAndServe()
	}
}

func (api *API) configureRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").HandlerFunc(HandleRoot).Methods("GET")
	router.HandleFunc("/auth", HandleAuth).Methods("POST")
	router.HandleFunc("/auth/verify", HandleAuthVerify).Methods("POST")
	//apiV1 := router.PathPrefix("/api/v1").Subrouter()

	return router
}
