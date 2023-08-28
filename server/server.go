package server

import "net/http"

// Server - структура сервера
type Server struct {
	router http.Handler
	port   string
}

// NewServer создаёт новый сервер
func NewServer(h http.Handler, port string) *Server {
	s := &Server{
		router: h,
		port:   port,
	}
	return s
}

// Run запускает сервер
func (s *Server) Run() error {
	return http.ListenAndServe(":"+s.port, s.router)
}
