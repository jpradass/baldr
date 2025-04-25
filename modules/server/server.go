package server

import (
	"fmt"
	"net/http"

	"github.com/jpradass/baldr/log"
)

type Server struct {
	port   int32
	host   string
	mux    *http.ServeMux
	logger log.Logger
	tls    bool
}

func New(host string, port int32, tls bool) *Server {
	logger, _ := log.New()
	return &Server{
		port:   port,
		host:   host,
		mux:    http.NewServeMux(),
		logger: logger,
		tls:    tls,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) Start() {
	s.logger.Infof("** starting server @ %s:%d **", s.host, s.port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.host, s.port), s); err != nil {
		s.logger.Error("there was an error starting the server. Error: %v", err)
	}
}
