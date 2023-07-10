package apiserver

import (
	"github.com/HellfastUSMC/goapi/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// APIServer API Server type
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New Create new instance of APIServer type
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start Starting API Server, return error if smth went wrong
func (s *APIServer) Start() error {
	if err := s.ConfLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	s.ConfRouter()
	if err := s.confStore(); err != nil {
		return err
	}
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

func (s *APIServer) confStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "hello")
		if err != nil {
			return
		}
	}
}
