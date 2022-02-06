package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/joshuaAllday/matillion/pkg/config"
	"gitlab.com/joshuaAllday/matillion/pkg/store"
	"gitlab.com/joshuaAllday/matillion/pkg/store/sqlstore"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/mtlog"
)

type Server struct {
	rootRouter *mux.Router
	config     *config.Config
	Store      store.Store
	server     *http.Server
	logger     *mtlog.Logger
}

func New(config *config.Config) (*Server, error) {
	s := &Server{
		rootRouter: mux.NewRouter(),
		config:     config,
	}
	s.initLogger()
	sqlstore, err := sqlstore.New(config.SqlSettings, config.SqlSettings.Seed)
	if err != nil {
		s.Logger().Error("unable to load store", mtlog.Fields{"err": err.Error()})
		return nil, err
	}
	s.Store = sqlstore

	return s, nil
}

func (s *Server) Start() error {
	handler := s.rootRouter
	s.server = &http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf(":%v", s.Config().Service.Port),
	}

	s.Logger().Info("starting server", mtlog.Fields{"port": s.Config().Service.Port})
	go s.server.ListenAndServe()
	return nil
}

func (s *Server) Stop() error {
	s.Logger().Info("stopping and shutting down server", mtlog.Fields{})
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func (s *Server) initLogger() {
	s.logger = mtlog.New(mtlog.DebugLevel)
}

func (s *Server) Router() *mux.Router {
	return s.rootRouter
}

func (s *Server) Config() *config.Config {
	return s.config
}

func (s *Server) Logger() *mtlog.Logger {
	return s.logger
}
