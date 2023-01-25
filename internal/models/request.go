package models

type CreateItemRequest struct {
	Name  string
	Price float32
}

type GetItemRequest struct {
	ID string
}
