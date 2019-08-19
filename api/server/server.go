package server

import (
	"net/http"
)

// HandlerFunc discare handler type
type HandlerFunc func(ctx Context, w http.ResponseWriter, r *http.Request)

// Server redirect request to handler
type Server struct {
	ctx  Context
	post *http.ServeMux
	get  *http.ServeMux
}

func New(ctx Context) *Server {
	return &Server{
		ctx:  ctx,
		post: http.NewServeMux(),
		get:  http.NewServeMux(),
	}
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.get.ServeHTTP(w, r)
	case http.MethodPost:
		s.post.ServeHTTP(w, r)
	}
}

// Post register post methods
func (s *Server) Post(pattern string, handler HandlerFunc) {
	s.post.HandleFunc(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(s.ctx, w, r)
	}))
}

// Get register get methods
func (s *Server) Get(pattern string, handler HandlerFunc) {
	s.get.HandleFunc(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(s.ctx, w, r)
	}))
}
