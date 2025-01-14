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
	"github.com/elc49/copod/cache"
	"github.com/elc49/copod/config"
	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/contracts"
	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/handlers"
	"github.com/elc49/copod/handlers/webhook"
	"github.com/elc49/copod/ip"
	copodMiddleware "github.com/elc49/copod/middleware"
	"github.com/elc49/copod/paystack"
	db "github.com/elc49/copod/sql"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/elc49/copod/tigris"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	r.Use(cors.AllowAll().Handler)
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
		r.Group(func(r chi.Router) {
			// Allow json content for below endpoints
			r.Use(middleware.AllowContentType("application/json"))
			r.With(copodMiddleware.Paystack).Handle("/webhook/paystack", webhook.Paystack())
			r.Handle("/ipinfo", handlers.Ipinfo())
			r.Handle("/early", handlers.EarlySignup())
		})
	})
	return r
}

func (s *Server) MountController() {
	// User
	uc := controller.User{}
	uc.Init(s.sql)
	// Title
	tc := controller.Title{}
	tc.Init(s.sql)
	// Support doc
	sc := controller.SupportingDoc{}
	sc.Init(s.sql)
	// Payment
	pc := controller.Payment{}
	pc.Init(s.sql)
	// Onboarding
	oc := controller.Onboarding{}
	oc.Init(s.sql)
	// Display picture
	dc := controller.DisplayPicture{}
	dc.Init(s.sql)
	// Early signup
	ec := controller.EarlySignup{}
	ec.Init(s.sql)
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

func (s *Server) CacheService() {
	cache.New()
}

func (s *Server) IpinfoService() {
	ip.New()
}

func (s *Server) NewEthereumService() {
	contracts.NewEthBackend()
}
