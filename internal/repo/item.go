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

func (i ItemsRepo) GetItemById(ctx context.Context, id string) (models.Item, error) {
	builder := sqlbuilder.NewSelectBuilder()
	sqlString, args := builder.
		Select("*").
		From("item").
		Where(builder.Equal("id", id)).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	row := i.conn.QueryRow(ctx, sqlString, args...)

	var item models.Item
	err := row.Scan(&item.ID, &item.Name, &item.ImageURL, &item.Count, &item.Price)
	return item, err
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
