package dto

type ReviewRequest struct {
	UserId int    `db:"user_id" json:"user_id"`
	Score  int    `db:"score" json:"score"`
	Text   string `db:"text" json:"text"`
	Pros   string `db:"pros" json:"pros"`
	Cons   string `db:"cons" json:"cons"`
}
