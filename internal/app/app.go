package app

import (
	"net/http"

	"github.com/proj-go-5/products/internal/storage"
)

type App struct {
	storage *storage.MySQLStorage
}

func NewApp(s *storage.MySQLStorage) *App {
	app := &App{
		storage: s,
	}
	return app
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
