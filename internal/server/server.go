package server

import (
    "net/http"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/rs/cors"
    "github.com/ikennarichard/contrib-tracker/internal/handler"
    "github.com/ikennarichard/contrib-tracker/internal/service"
)

type Server struct {
    httpServer *http.Server
}

func NewServer(svc *service.ContributionService, port string) *Server {
    r := chi.NewRouter()

    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.Timeout(60 * time.Second))

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
    })

    r.Use(c.Handler)

    contributionHandler := handler.NewContributionHandler(svc)
    contributionHandler.RegisterRoutes(r)

    return &Server{
        httpServer: &http.Server{
            Addr:    ":" + port,
            Handler: r,
        },
    }
}

func (s *Server) Start() error {
    return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
    return s.httpServer.Shutdown(nil)
}