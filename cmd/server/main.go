package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/proj-go-5/products/internal/app"
	"github.com/proj-go-5/products/internal/storage"
)

var (
	Version = ""
)

func main() {
	fmt.Printf("App version : %s\n", Version)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		fmt.Println("Port not found")
		port = "8000"
	} else {
		fmt.Printf("Running the server on port %s\n", port)
	}

	mysql := storage.NewStorage()
	defer mysql.Close()

	app := app.NewApp(mysql)
	router := app.GetRouter()

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		fmt.Print("Server run error", err)
	}
}
