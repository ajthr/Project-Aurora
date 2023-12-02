package main

import (
	"log"
	"net/http"

	"identity-service/internal/routers"
)

func main() {
	router := routers.NewRouter()

	log.Println("Ready to receive requests.")
	http.ListenAndServe(":7000", router)
}
