package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/elc49/copod/config"
	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/handlers"
	"github.com/elc49/copod/handlers/webhook"
	copodMiddleware "github.com/elc49/copod/middleware"
	"github.com/elc49/copod/paystack"
	db "github.com/elc49/copod/sql"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/elc49/copod/tigris"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	sql *sql.Queries
}

func New() *Server {
	// Setup config variables
	config.New()

	s := &Server{}
	return s
}

func (s *Server) Start() {
	server := &http.Server{Addr: "0.0.0.0:" + config.C.Server.Port, Handler: s.MountRouter()}
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
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
	// Wait for server ctx to be stopped
	<-sCtx.Done()
}

func (s *Server) MountRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.CleanPath)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	r.Handle("/graphql", handlers.GraphQL())
	r.Route("/api", func(r chi.Router) {
		r.Handle("/upload", handlers.UploadDoc())
		r.With(copodMiddleware.Paystack).Handle("/webhook/paystack", webhook.Paystack())
	})
	return r
}

func (s *Server) MountController() {
	// User
	u := controller.User{}
	u.Init(s.sql)
	// Upload
	p := controller.Upload{}
	p.Init(s.sql)
}

func (s *Server) Database(opt postgres.Postgres) {
	sqlStore := db.InitDB(config.C.Database.Rdbms)
	s.sql = sqlStore
}

func (s *Server) TigrisService() {
	tigris.New()
}

func (s *Server) PaystackService() {
	paystack.New(s.sql)
}
