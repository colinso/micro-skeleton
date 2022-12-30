package repo

import (
	"Personal/micro-skeleton/internal/models"
	"context"

	"github.com/apex/log"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
)

type ItemsRepo struct {
	conn *pgx.Conn
}

func NewItemsRepo(conn *pgx.Conn) *ItemsRepo {
	return &ItemsRepo{
		conn: conn,
	}
}

func (i ItemsRepo) GetItemById() models.Item {
	return models.Item{}
}

func (i ItemsRepo) CreateItem(item models.Item) error {
	builder := sqlbuilder.NewInsertBuilder()
	sqlString, args := builder.
		InsertInto("item").
		Cols("id", "name", "image_url", "count", "price").
		Values(item.ID, item.Name, item.ImageURL, item.Count, item.Price).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := i.conn.Exec(context.Background(), sqlString, args...)
	if err != nil {
		log.WithError(err).Error("Error creating item in DB")
		return err
	}
	return nil
}

func (i ItemsRepo) UpdateItem(item models.Item) models.Item {
	return models.Item{}
}
