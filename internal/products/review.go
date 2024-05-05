package products

type Review struct {
	Id        int    `db:"id"`
	ProductId int    `db:"product_id"`
	UserId    int    `db:"user_id"`
	Score     int    `db:"score"`
	Text      string `db:"text"`
	Pros      string `db:"pros"`
	Cons      string `db:"cons"`
}
