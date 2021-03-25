package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"

	middleware "github.com/aacebo/equinox-api/src/middleware"
)

func Router() *chi.Mux {
	router := chi.NewRouter()

	router.Use(chi_middleware.Logger)
	router.Use(chi_middleware.Recoverer)

	router.With(middleware.WithParams(struct {
	}{})).Get("/{oslug:[a-z-]+}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	return router
}
