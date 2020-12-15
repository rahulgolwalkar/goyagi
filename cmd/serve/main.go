// cmd/serve/main.go
package main

import (
	"github.com/rahulgolwalkar/goyagi/pkg/application"
	"net/http"

	"github.com/rahulgolwalkar/goyagi/pkg/server"
	"github.com/lob/logger-go"
)

func main() {
	log := logger.New()

	app, err1 := application.New()

	if err1 != nil {
		log.Err(err1).Fatal("failed to initialize application")
	}

	srv := server.New(app)

	log.Info("server started", logger.Data{"port": app.Config.Port})

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Err(err).Fatal("serve stopped")
	}

	log.Info("serve stopped")
}