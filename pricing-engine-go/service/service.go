package service

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"

	"pricingengine/service/app"
	"pricingengine/service/rpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Start begins a chi-Mux'd net/http server on port 3000
func Start() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(5 * time.Second))

	rpc := rpc.RPC{
		App: &app.App{},
	}
	log.Info("server started")
	r.Post("/generate_pricing", rpc.GeneratePricing)
	http.ListenAndServe(":3000", r)
}
