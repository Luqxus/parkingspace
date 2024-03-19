package main


import (
	"github.com/luquxSentinel/spacedrive/service"
	"net/http"
)

type APIServer struct {
	mux *http.ServeMux
	listenAddress string
	authService service.AuthService
}

func NewAPIServer(listenAddress string, authService service.AuthService) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		mux: http.NewServeMux(),
	}
}

func (api *APIServer) Run() error {
	
	return http.ListenAndServe(api.listenAddress, api.mux)
}


