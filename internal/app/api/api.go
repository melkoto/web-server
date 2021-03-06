package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// API is base server instance description
type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New is API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start is http server/configure loggers, router, db connection ...
func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port:", api.config.BindAddr)
	api.configureRouterField()
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
