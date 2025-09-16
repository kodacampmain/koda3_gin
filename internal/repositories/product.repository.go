package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda3_gin/internal/models"
	"github.com/redis/go-redis/v9"
)

type ProductRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewProductRepository(db *pgxpool.Pool, rdb *redis.Client) *ProductRepository {
	return &ProductRepository{
		db:  db,
		rdb: rdb,
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

func (p *ProductRepository) GetProducts(rctx context.Context) ([]models.ProductData, error) {
	// cache-aside pattern
	// cek redis terlebih dahulu
	redisKey := "fakhri:products-all"
	cmd := p.rdb.Get(rctx, redisKey)
	if cmd.Err() != nil {
		if cmd.Err() == redis.Nil {
			log.Printf("Key %s does not exist\n", redisKey)
		} else {
			log.Println("Redis Error. \nCause: ", cmd.Err().Error())
		}
	} else {
		// cache hit
		var cachedProducts []models.ProductData
		cmdByte, err := cmd.Bytes()
		if err != nil {
			log.Println("Internal Server Error.\nCause: ", err.Error())
		} else {
			if err := json.Unmarshal(cmdByte, &cachedProducts); err != nil {
				log.Println("Internal Server Error.\nCause: ", err.Error())
			}
			if len(cachedProducts) > 0 {
				return cachedProducts, nil
			}
		}
	}

	// cache miss
	sql := "SELECT id, name, promo_id, price, created_at, updated_at FROM products"
	var products []models.ProductData
	rows, err := p.db.Query(rctx, sql)
	if err != nil {
		return []models.ProductData{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var product models.ProductData
		if err := rows.Scan(&product.Id, &product.Name, &product.PromoId, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return []models.ProductData{}, err
		}
		// log.Println(product)
		products = append(products, product)
		// log.Println(products)
	}

	// renew cache
	bt, err := json.Marshal(products)
	if err != nil {
		log.Println("Internal Server Error.\nCause: ", err.Error())
	} else {
		if err := p.rdb.Set(rctx, redisKey, string(bt), 5*time.Minute).Err(); err != nil {
			log.Println("Redis Error.\nCause: ", err.Error())
		}
	}

	return products, nil
}
