package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) route() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/terminal", app.VirtualTerminal)
	return mux
}
