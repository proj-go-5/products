package dto

type ProductRequest struct {
	Title       string `json:"title"`
	Price       int32  `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image_url"`
}
