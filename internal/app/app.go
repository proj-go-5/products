package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/proj-go-5/products/internal/dto"
	"github.com/proj-go-5/products/internal/storage"
	"github.com/proj-go-5/products/internal/utils"
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
	router.HandleFunc("POST /products", a.handleAddProduct)

	return router
}

func (a *App) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func (a *App) handleAddProduct(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := utils.ReadBody(r)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	pr := dto.ProductRequest{}
	err = json.Unmarshal(bodyBytes, &pr)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Body unmarshalling error: %v", err))
		return
	}

	t := time.Now()
	err = a.storage.Add(
		map[string]interface{}{
			"title":       pr.Title,
			"price":       pr.Price,
			"description": pr.Description,
			"update_date": t.Format("2006-01-02 15:04:05"),
			"images":      pr.Image,
		},
		"Product",
	)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Insert error: %v", err))
		return
	}
	sendOk(w)
}

func sendError(w http.ResponseWriter, status int, text string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`{"status":"error","message":%s"}`, text)))
}

func sendOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
