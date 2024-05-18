package dto

import "time"

type ProductRequest struct {
	Id          int32     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Price       int32     `db:"price" json:"price"`
	Description string    `db:"description" json:"description"`
	Image       string    `db:"image" json:"image_url"`
	UpdateDate  time.Time `db:"-"`
}
