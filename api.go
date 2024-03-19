package main


import (
	"context"
	"debug/macho"
	"github.com/luquxSentinel/spacedrive/service"
	"net/http"
	"time"
)

type APIFunc func(w http.ResponseWriter, r *http.Request)

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

func APIFuncHandler(fn APIFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// context with timeout from request context
		ctx, cancel := context.WithTimeout(request.Context(), 30 * time.Second)
		defer cancel()

		// handle APIFunction
		fn(writer, request.WithContext(ctx))
	}
}


