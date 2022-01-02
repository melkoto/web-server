package api

import "github.com/sirupsen/logrus"

// configureLoggerField used to configure API instance
func (api *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}

	api.logger.SetLevel(logLevel)
	return nil
}
