package monito

import (
	"log"
	"net/http"

	"github.com/PI-Victor/monito/pkg/api"
)

const (
	apiv1   = "/api/v1"
	rooturl = "/"
)

func initMuxServerAndRoutes(s *Server) *http.ServeMux {
	newMuxServer := http.NewServeMux()
	newMuxServer.Handle(rooturl, s.ListAPIEndpoints)
	namespaceRoute := &api.Namespace{}
	return newMuxServer
}

// InitializeAPI starts a goroutine with a simple http server
func InitializeAPI(s *Server) {
	newMuxServerHandler := initMuxServerAndRoutes(s)
	server := &http.Server{
		Addr: "127.0.0.1:8080", //&Server.bindAddress,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}
