package main

import (
	"log"
	"mongoservice/router"
	"net/http"
)

func main() {
	routers := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8888", routers))
}
