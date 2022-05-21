package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/mendezdev/go-postgres-demo/pkg/db"
)

func NewAPI(pgdb *pg.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger, jsonContentMiddleware, middleware.WithValue("DB", pgdb))

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

func jsonContentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

type CreateHomeRequest struct {
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	AgentID     int64  `json:"agent_id"`
}

type CreateHomeResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Home    *db.Home `json:"home"`
}

func createHome(w http.ResponseWriter, r *http.Request) {
	req := &CreateHomeRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &CreateHomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &CreateHomeResponse{
			Success: false,
			Error:   "could not get database from context",
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	home, err := db.CreateHome(pgdb, &db.Home{
		Price:       req.Price,
		Description: req.Description,
		Address:     req.Address,
		AgentID:     req.AgentID,
	})

	if err != nil {
		res := &CreateHomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	res := &CreateHomeResponse{
		Success: true,
		Error:   "",
		Home:    home,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

type GetHomesResponse struct {
	Homes []db.Home `json:"homes"`
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

type DeleteHomeByIDResponse struct {
	Success bool `json:"success"`
}

func deleteHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")
	w.Write([]byte(fmt.Sprintf("delete home by ID: %s", homeID)))
}
