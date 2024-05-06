package products

import "time"

type Product struct {
	Id          int       `db:"id"`
	Title       string    `db:"title"`
	Price       float64   `db:"price"`
	Description string    `db:"description"`
	UpdatedAt   time.Time `db:"updated_date"`
	ImageUrl    string    `db:"image_url"`
}
