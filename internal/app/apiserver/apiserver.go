package apiserver

//API Server type
type APIServer struct {
	config *Config
}

//Create new instance of APIServer type
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

//Starting API Server, return error if smth went wrong
func (s *APIServer) Start() error {
	return nil
}
