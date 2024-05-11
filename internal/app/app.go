package app

import "net/http"

type App struct {
	// Storage
}

func (a *App) GetRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /products", a.handleGetProducts)
	// handle other functions

	return router
}

func (a *App) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
