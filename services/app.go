package main

import (
	"fmt"
	"os"
	"services/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	server.Run()
	fmt.Println("Service started...")
}
