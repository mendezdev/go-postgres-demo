package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mendezdev/go-postgres-demo/pkg/api"
)

func main() {
	router := api.NewAPI()

	log.Print("we are up and running!")
	port := os.Getenv("PORT")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
