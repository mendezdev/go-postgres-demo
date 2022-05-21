package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mendezdev/go-postgres-demo/pkg/api"
	"github.com/mendezdev/go-postgres-demo/pkg/db"
)

func main() {
	pgdb, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	router := api.NewAPI(pgdb)

	log.Print("we are up and running!")
	port := os.Getenv("PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
