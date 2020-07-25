package main

import (
	"log"
	"net/http"
	"services/router"
)

func main() {
	routers := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8888", routers))
}
