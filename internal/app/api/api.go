package api

import "github.com/sirupsen/logrus"

// API is base server instance description
type API struct {
	config *Config
	logger *logrus.Logger
}

// New is API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
	}
}

// Start is http server/configure loggers, router, db connection ...
func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port:", api.config.BindAddr)
	return nil
}
