package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elc49/copod/handlers"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct{}

func New() *Server {
	s := &Server{}
	return s
}

func (s *Server) Start() {
	server := &http.Server{Addr: "0.0.0.0:4545", Handler: s.mount()}
	// Server ctx
	sCtx, sStopCtx := context.WithCancel(context.Background())
	// Listen for syscall signals(interrupt/quit)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		<-sig
		// 30 seconds grace period
		shutCtx, _ := context.WithTimeout(sCtx, time.Second*30)
		go func() {
			<-shutCtx.Done()
			if shutCtx.Err() == context.DeadlineExceeded {
				log.Fatalln("shutdown grace period tiemout")
			}
		}()
		// Trigger shutdown
		err := server.Shutdown(shutCtx)
		if err != nil {
			log.Fatalln(err)
		}
		sStopCtx()
	}()
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
	// Wait for server ctx to be stopped
	<-sCtx.Done()
}

func (s *Server) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.CleanPath)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	r.Handle("/graphql", handlers.GraphQL())
	return r
}

func main() {
	s := New()
	s.Start()
}
