package main

import (
	"flag"
	"fmt"
	"net/http"
	"products/internal/app"
)

func main() {
	portFlag := flag.Int("port", 8000, "A port to run the service on")
	flag.Parse()

	app := app.App{}
	router := app.GetRouter()

	if portFlag != nil {
		fmt.Println("PORT!", *portFlag)
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), router)
	if err != nil {
		fmt.Print("Server run error", err)
	}
}
