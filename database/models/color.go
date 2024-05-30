package models

type Color struct {
	ID              string `db:"id" json:"id"`
	Title           string `db:"title" json:"title"`
	InstagramImages string `db:"instagram_images" json:"instagram_images"`
	WebsiteImages   string `db:"website_images" json:"website_images"`
	Videos          string `db:"videos" json:"videos"`
}
