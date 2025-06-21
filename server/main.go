package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file\n%s", err.Error())
		return
	}

	port := os.Getenv("PORT")

	fmt.Printf("Listening at port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Printf("Error Listening \n %s", err.Error())
	}
}
