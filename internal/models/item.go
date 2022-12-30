package models

type Item struct {
	ID       string  `json:"item_id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Count    int     `json:"count"`
	ImageURL string  `json:"image_url"`
}
