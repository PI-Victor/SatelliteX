package monito

import (
	"fmt"
	"net/http"
	"path"

	"github.com/PI-Victor/monito/pkg/log"
)

const (
	apiv1       = "/api/v1"
	metricspath = "/metrics"
	rooturl     = "/"
)

func loadMuxRoutes(s *Server) *http.ServeMux {
	newMuxServer := http.NewServeMux()
	newMuxServer.HandleFunc(rooturl, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a test!!!"))
	})
	return newMuxServer
}

func MetricsUrlPath(w http.ResponseWriter, r *http.Request) {

	fmt.Println("this is a test")
}

func loadTestAPIRoutes(s *http.ServeMux) *http.ServeMux {
	s.HandleFunc(metricspath, MetricsUrlPath)
	return s
}

// InitializeAPI starts a goroutine with a simple http server
func InitializeAPI(s *Server) {
	newMuxServerHandler := loadMuxRoutes(s)
	server := &http.Server{
		Addr:    "127.0.0.1:8080", //&Server.bindAddress,
		Handler: newMuxServerHandler,
	}
	log.Info("API Available at http://%s%s", server.Addr, apiv1)
	go func() {
		log.Panic("The Api Server encountered an error %q", server.ListenAndServe())
	}()
}

// a simple function for testing url formatting
func testformatURLPaths(urlPath string) string {
	// just a test url formatter
	formattedUrlPath := path.Join(apiv1, urlPath)
	return formattedUrlPath
}
