package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = "28082"
)

// getEnv получает значение переменной окружения или возвращает значение по умолчанию.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// newReverseProxy создает новый обработчик ReverseProxy с указанной целевой схемой и хостом.
func newReverseProxy(targetScheme, targetHost string) http.Handler {
	targetURL := &url.URL{
		Scheme: targetScheme,
		Host:   targetHost,
	}

	// Проверяем, что целевой хост указан
	if targetHost == "" {
		log.Fatal("[ERROR] Target host is not specified. Set the PROXY_TARGET_HOST environment variable.")
	}
	// Проверяем, что целевая схема указана
	if targetScheme == "" {
		log.Fatal("[ERROR] Target scheme is not specified. Set the PROXY_TARGET_SCHEME environment variable.")
	}

	log.Printf("[INFO] Configuring reverse proxy to target: %s", targetURL.String())

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Оригинальный Director для NewSingleHostReverseProxy делает примерно то же самое,
	// но мы можем добавить свою логику при необходимости.
	// Здесь мы просто устанавливаем заголовок Host.
	defaultDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		defaultDirector(req)
		// Устанавливаем заголовок Host равным целевому хосту,
		// чтобы целевой сервер правильно обрабатывал запросы (особенно важно для виртуальных хостов).
		req.Host = targetHost
		log.Printf("[DEBUG] Proxying request for %s to %s, path: %s", req.RemoteAddr, targetURL, req.URL.Path)
		log.Printf("[DEBUG] Request Headers: %v", req.Header)
	}

	// ModifyResponse используется для логирования заголовков ответа от целевого сервера.
	proxy.ModifyResponse = func(resp *http.Response) error {
		log.Printf("[DEBUG] Received response from target %s for %s. Status: %s", targetURL, resp.Request.RemoteAddr, resp.Status)
		log.Printf("[DEBUG] Response Headers from target: %v", resp.Header)
		// Мы можем изменять ответ здесь при необходимости, например, добавлять заголовки.
		// resp.Header.Set("X-Proxied-By", "Awesome-Go-Proxy")
		return nil // Возвращаем nil, если не было ошибок при модификации
	}

	// ErrorHandler логирует ошибки, возникшие во время проксирования.
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		log.Printf("[ERROR] Proxy error: %v", err)
		// Отправляем стандартный ответ "Bad Gateway", если произошла ошибка связи с целевым сервером
		rw.WriteHeader(http.StatusBadGateway)
		// Можно отправить и тело ответа с описанием ошибки, но будьте осторожны, чтобы не раскрыть лишнюю информацию.
		// _, _ = rw.Write([]byte("Proxy error: " + err.Error()))
	}

	return proxy
}

func main() {
	// --- Конфигурация ---
	listenHost := getEnv("PROXY_LISTEN_HOST", defaultHost)
	listenPort := getEnv("PROXY_LISTEN_PORT", defaultPort)
	certFile := getEnv("PROXY_TLS_CERT_PATH", "")          // Пример: PROXY_TLS_CERT_PATH=/path/to/cert.pem
	keyFile := getEnv("PROXY_TLS_KEY_PATH", "")            // Пример: PROXY_TLS_KEY_PATH=/path/to/key.pem
	targetScheme := getEnv("PROXY_TARGET_SCHEME", "https") // Пример: PROXY_TARGET_SCHEME=https
	targetHost := getEnv("PROXY_TARGET_HOST", "")          // Пример: PROXY_TARGET_HOST=api.example.com

	// Проверка обязательных параметров
	if targetHost == "" {
		log.Fatal("[ERROR] Required environment variable PROXY_TARGET_HOST is not set.")
	}
	// Дополнительная проверка схемы для ясности
	if targetScheme != "http" && targetScheme != "https" {
		log.Fatalf("[ERROR] Invalid PROXY_TARGET_SCHEME: %s. Must be 'http' or 'https'.", targetScheme)
	}

	log.Printf("[INFO] Process PID: %d, PPID: %d", os.Getpid(), os.Getppid())

	// --- Создание обработчика ---
	proxyHandler := newReverseProxy(targetScheme, targetHost)

	// --- Настройка сервера ---
	servingAddr := fmt.Sprintf("%s:%s", listenHost, listenPort)
	server := &http.Server{
		Addr:    servingAddr,
		Handler: proxyHandler,
		// Можно добавить таймауты для повышения стабильности
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
		// IdleTimeout:  120 * time.Second,
	}

	// --- Запуск сервера ---
	go func() {
		if certFile != "" && keyFile != "" {
			log.Printf("[INFO] Starting HTTPS server at https://%s", servingAddr)
			log.Printf("[INFO] Using TLS certificate: %s", certFile)
			log.Printf("[INFO] Using TLS key: %s", keyFile)
			if err := server.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
				log.Fatalf("[ERROR] Failed to start TLS server: %v", err)
			}
		} else {
			log.Printf("[INFO] Starting HTTP server at http://%s", servingAddr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("[ERROR] Failed to start server: %v", err)
			}
		}
	}() // Запускаем сервер в отдельной горутине

	// --- Graceful Shutdown ---
	quit := make(chan os.Signal, 1) // Канал для сигналов ОС
	// Ожидаем сигналы SIGINT (Ctrl+C) и SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("[INFO] Received signal %v. Shutting down server...", <-quit)

	// Создаем контекст с таймаутом для завершения работы
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Даем 30 секунд на завершение запросов
	defer cancel()

	// Вызываем Shutdown для плавного завершения работы сервера
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("[ERROR] Server shutdown failed: %v", err)
	}

	log.Println("[INFO] Server gracefully stopped")
}
