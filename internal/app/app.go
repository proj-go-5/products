package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	router.HandleFunc("PUT /products", a.handleChangeProduct)
	router.HandleFunc("DELETE /products/{id}", a.handleDeleteProduct)
	router.HandleFunc("GET /products/{id}", a.handleGetProduct)

	router.HandleFunc("GET /products/{id}/reviews", a.handleGetReview)
	router.HandleFunc("POST /products/{id}/reviews", a.handleAddReview)
	router.HandleFunc("DELETE /products/{id}/reviews/{review_id}", a.handleDeleteReview)

	return router
}

func (a *App) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")
	pageNum := r.URL.Query().Get("page_num")
	pageSize := r.URL.Query().Get("page_size")
	ids := r.URL.Query().Get("ids")

	pageNumInt, err := strconv.Atoi(pageNum)
	if err != nil || pageNumInt < 1 {
		pageNumInt = 1
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		pageSizeInt = 100
	}

	products, err := a.storage.GetProducts(filter, sortBy, order, pageNumInt, pageSizeInt, ids)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// In the future transfer products to the API
	fmt.Println(products)

	sendOk(w)
}

func (a *App) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	productIdStr := r.PathValue("id")
	productId, err := strconv.ParseInt(productIdStr, 10, 32)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Cannot parse id: %v", err))
		return
	}

	var products []dto.ProductRequest
	err = a.storage.Get(&products, "Product", fmt.Sprintf("id = %d", productId))
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendOk(w)
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

func (a *App) handleChangeProduct(w http.ResponseWriter, r *http.Request) {
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
	err = a.storage.UpdateProduct(&pr)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Update error: %v", err))
		return
	}
	sendOk(w)
}

func (a *App) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	productIdStr := r.PathValue("id")
	productId, err := strconv.ParseInt(productIdStr, 10, 32)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Cannot parse id: %v", err))
		return
	}
	err = a.storage.Delete("Product", int32(productId))
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Update error: %v", err))
		return
	}
	sendOk(w)
}

func (a *App) handleGetReview(w http.ResponseWriter, r *http.Request) {
	productIdStr := r.PathValue("id")
	productId, err := strconv.ParseInt(productIdStr, 10, 32)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Cannot parse id: %v", err))
		return
	}

	var reviews []dto.ProductRequest
	err = a.storage.Get(&reviews, "Review", fmt.Sprintf("product_id = %d", productId))
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendOk(w)
}

func (a *App) handleAddReview(w http.ResponseWriter, r *http.Request) {
}

func (a *App) handleDeleteReview(w http.ResponseWriter, r *http.Request) {
}

func sendError(w http.ResponseWriter, status int, text string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`{"status":"error","message":%s"}`, text)))
}

func sendOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
