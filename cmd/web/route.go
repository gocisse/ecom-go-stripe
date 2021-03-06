package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) route() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/terminal", app.VirtualTerminal)
	mux.Post("/payment-succeeded", app.succeededPayment)

	mux.Get("/charge-once", app.ChargeOnce)
	
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*" , http.StripPrefix("/static", fileServer))

	return mux
}
