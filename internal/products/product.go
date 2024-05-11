package products

import "time"

type Product struct {
	Id          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Price       float64   `db:"price" json:"price"`
	Description string    `db:"description" json:"description"`
	UpdatedAt   time.Time `db:"updated_date" json:"updated_date"`
	ImageUrl    string    `db:"image_url" json:"image_url"`
}
