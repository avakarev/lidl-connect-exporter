// Package metrics implements prometheus metrics
package metrics

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// DefaultPort is default http port
const DefaultPort = "9100"

// DefaultPath is default metrics route path
const DefaultPath = "/metrics"

var port string
var path string

// Serve starts an HTTP server
func Serve() error {
	http.Handle(path, promhttp.Handler())
	server := &http.Server{
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         fmt.Sprintf(":%s", port),
	}
	log.Info().Str("port", port).Str("path", path).Msg("start serving metrics")
	return server.ListenAndServe()
}

func init() {
	port = DefaultPort
	if p := os.Getenv("HTTP_PORT"); p != "" {
		port = p
	}

	path = DefaultPath
	if p := os.Getenv("METRICS_PATH"); p != "" {
		path = p
	}
}
