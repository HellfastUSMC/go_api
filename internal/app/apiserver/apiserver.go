package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// API Server type
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// Create new instance of APIServer type
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Starting API Server, return error if smth went wrong
func (s *APIServer) Start() error {
	if err := s.ConfLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	s.ConfRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) ConfLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLVL)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) ConfRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
