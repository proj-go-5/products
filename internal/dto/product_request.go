package dto

import "time"

type ProductRequest struct {
	Id          int32     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Price       int32     `db:"price" json:"price"`
	Description string    `db:"description" json:"description"`
	UpdateDate  time.Time `db:"update_date" json:"update_date"`
	Image       string    `db:"images" json:"image_url"`
}

type ProductPageRequest struct {
	Page []ProductRequest `json:"page"`
}
