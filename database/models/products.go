package models

type Product struct {
	ID              string   `db:"id" json:"id"`
	Title           string   `db:"title" json:"title"`
	Description     string   `db:"description" json:"description"`
	Price           string   `db:"price" json:"price"`
	Code            string   `db:"code" json:"code"`
	InstagramImages []string `db:"instagram_images" json:"instagram_images"`
	WebsiteImages   []string `db:"website_images" json:"website_images"`
	Videos          []string `db:"videos" json:"videos"`
}
