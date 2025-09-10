package repositories

import (
	"context"
	"fmt"
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

func (p *ProductRepository) AddNewProduct(rctx context.Context, body models.Product) (models.ProductData, error) {
	sql := "INSERT INTO products (name, promo_id, price) VALUES ($1,$2,$3) RETURNING id, name"
	values := []any{body.Name, body.PromoId, body.Price}
	var newProduct models.ProductData
	if err := p.db.QueryRow(rctx, sql, values...).Scan(&newProduct.Id, &newProduct.Name); err != nil {
		log.Println("Internal Server Error: ", err.Error())
		return models.ProductData{}, err
	}
	return newProduct, nil
}

func (p *ProductRepository) InsertNewProduct(rctx context.Context, body models.Product) (pgconn.CommandTag, error) {
	sql := "INSERT INTO products (name, promo_id, price) VALUES ($1,$2,$3)"
	values := []any{body.Name, body.PromoId, body.Price}
	return p.db.Exec(rctx, sql, values...)
}

// func (p *ProductRepository) Add() {}

func (p *ProductRepository) EditProduct(rctx context.Context, body models.EditProductBody, id int) (models.ProductData, error) {
	sql := "UPDATE products SET "
	values := []any{}

	if body.Name != nil {
		sql += fmt.Sprintf("%s=$%d, ", "name", len(values)+1)
		values = append(values, body.Name)
	}

	if body.PromoId != nil {
		sql += fmt.Sprintf("%s=$%d, ", "promo_id", len(values)+1)
		values = append(values, body.PromoId)
	}

	if body.Price != nil {
		sql += fmt.Sprintf("%s=$%d, ", "price", len(values)+1)
		values = append(values, body.Price)
	}

	sql += fmt.Sprintf("updated_at=now() WHERE id=$%d RETURNING id, name, price, promo_id, created_at, updated_at", len(values)+1)
	values = append(values, id)

	var product models.ProductData
	if err := p.db.QueryRow(rctx, sql, values...).Scan(&product.Id, &product.Name, &product.Price, &product.PromoId, &product.CreatedAt, &product.UpdatedAt); err != nil {
		return models.ProductData{}, err
	}
	return product, nil
	// return sql, values
}
