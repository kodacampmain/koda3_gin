package repositories

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda3_gin/internal/models"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) AddNewProduct(rctx context.Context, body models.Product) (models.Product, error) {
	sql := "INSERT INTO products (name, promo_id, price) VALUES ($1,$2,$3) RETURNING id, name"
	values := []any{body.Name, body.PromoId, body.Price}
	var newProduct models.Product
	if err := p.db.QueryRow(rctx, sql, values...).Scan(&newProduct.Id, &newProduct.Name); err != nil {
		log.Println("Internal Server Error: ", err.Error())
		return models.Product{}, err
	}
	return newProduct, nil
}

func (p *ProductRepository) InsertNewProduct(rctx context.Context, body models.Product) (pgconn.CommandTag, error) {
	sql := "INSERT INTO products (name, promo_id, price) VALUES ($1,$2,$3)"
	values := []any{body.Name, body.PromoId, body.Price}
	return p.db.Exec(rctx, sql, values...)
}
