package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test/internal/http/handlers"
	"test/pkg/middlewares"
	"time"
)

func Run() {
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	r.Handle("/ports", handlers.NewPortsHandler()).Methods(http.MethodPost)
	r.Use(middlewares.PrometheusMiddleware, middlewares.RequestLogger)
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       1 * time.Hour,
		ReadHeaderTimeout: 2 * time.Hour,
		WriteTimeout:      1 * time.Hour,
		IdleTimeout:       30 * time.Hour,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Server Stopped: %#v", err)
		}
	}()

	serverError := make(chan error, 1)
	gracefulShutdown := make(chan os.Signal)
	signal.Notify(gracefulShutdown, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	log.Print("Server Started")
	select {
	case serverErr := <-serverError:
		log.Printf("server error : %#v", serverErr)
		shutdown(srv)
	case sig := <-gracefulShutdown:
		log.Printf("process canceled by signal : %s", sig.String())
		shutdown(srv)
	}

	log.Print("Server Exited Properly")

}

func shutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}
		cancel()
	}()
}
