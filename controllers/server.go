package controllers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hud0shnik/vallhallaapi/internal/handler"
)

// Server - структура сервера
type Server struct {
	basePath       string
	requestTimeout time.Duration
	router         http.Handler
	Server         *http.Server
}

// NewServer создаёт новый сервер
func NewServer(config *Config) *Server {

	s := &Server{
		basePath:       config.BasePath,
		requestTimeout: config.RequestTimeout,
	}

	s.NewRouter()

	s.Server = &http.Server{
		Addr:              config.ServerPort,
		Handler:           s.router,
		ReadTimeout:       config.RequestTimeout,
		ReadHeaderTimeout: config.RequestTimeout,
	}

	return s
}

// NewRouter создаёт новый роутер
func (s *Server) NewRouter() {

	// Роутер
	router := chi.NewRouter()
	router.Use(middleware.Timeout(s.requestTimeout))

	// Маршруты
	router.Get(s.basePath+"/search", handler.Search)
	router.Get(s.basePath+"/info", handler.Info)

	s.router = router

}
