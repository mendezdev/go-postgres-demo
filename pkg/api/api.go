package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewAPI() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/homes", func(r chi.Router) {
		r.Post("/", createHome)
		r.Get("/{homeID}", getHomeByID)
		r.Get("/", getHomes)
		r.Put("/{homeID}", updateHomeByID)
		r.Delete("/{homeID}", deleteHomeByID)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	return r
}

func createHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create homes"))
}

func getHomes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get homes"))
}

func getHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")
	w.Write([]byte(fmt.Sprintf("get home by ID: %s", homeID)))
}

func updateHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")
	w.Write([]byte(fmt.Sprintf("update home by ID: %s", homeID)))
}

func deleteHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")
	w.Write([]byte(fmt.Sprintf("delete home by ID: %s", homeID)))
}
