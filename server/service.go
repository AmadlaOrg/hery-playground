package server

// NewServerService to set up the server service
func NewServerService() IServer {
	return &SServer{}
}
