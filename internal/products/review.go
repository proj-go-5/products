package products

type Review struct {
	Id        int    `db:"id" json:"id"`
	ProductId int    `db:"product_id" json:"product_id"`
	UserId    int    `db:"user_id" json:"user_id"`
	Score     int    `db:"score" json:"score"`
	Text      string `db:"text" json:"text"`
	Pros      string `db:"pros" json:"pros"`
	Cons      string `db:"cons" json:"cons"`
}
