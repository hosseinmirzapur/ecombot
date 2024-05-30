package models

type Product struct {
	ID          string `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Price       string `db:"price" json:"price"`
	Code        string `db:"code" json:"code"`
	Colors      string `db:"colors" json:"colors"`
}
