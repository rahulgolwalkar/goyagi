// cmd/serve/main.go
package main

import (
	"net/http"

	"github.com/rahulgolwalkar/goyagi/pkg/server"
	"github.com/lob/logger-go"
)

func main() {
	log := logger.New()

	srv := server.New()

	log.Info("serve started")

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Err(err).Fatal("serve stopped")
	}

	log.Info("serve stopped")
}