package models

type Item struct {
	ID       string  `json:"item_id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Price    float32 `json:"price" db:"price"`
	Count    int     `json:"count" db:"count"`
	ImageURL string  `json:"image_url" db:"image_url"`
}
