package api

// API is base server instance description
type API struct {
	config *Config
}

// New is API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

// Start is http server/configure loggers, router, db connection ...
func (api *API) Start() error {
	return nil
}
