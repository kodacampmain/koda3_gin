package models

import "time"

type Product struct {
	Name    string `json:"name" binding:"required"`
	PromoId *int   `json:"promo_id"`
	Price   int    `json:"price" binding:"required"`
}

type EditProductBody struct {
	PromoId *int    `json:"promo_id"`
	Name    *string `json:"name"`
	Price   *int    `json:"price"`
}

type ProductResponse struct {
	SuccessResponse
	Data ProductData `json:"data"`
}

type ProductData struct {
	Id        int        `db:"id"`
	Name      string     `db:"name"`
	PromoId   *int       `db:"promo_id"`
	Price     int        `db:"price"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
