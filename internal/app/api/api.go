package api

// API is base server instance description
type API struct {
}

// New is API constructor: build base API instance
func New() *API {
	return &API{}
}

// Start is http server/configure loggers, router, db connection ...
func (api *API) Start() error {
	return nil
}
