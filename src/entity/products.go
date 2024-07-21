package entity

import "time"

type Products struct {
	ID            int64     `json:"id" db:"id,omitempty" param:"id"`
	CategoryID    int64     `json:"category_id" db:"category_id,omitempty" param:"category_id"`
	Name          string    `json:"name" db:"name,omitempty" param:"name"`
	Description   string    `json:"description" db:"description,omitempty" param:"description"`
	Price         float64   `json:"price" db:"price,omitempty" param:"price"`
	DiscountPrice float64   `json:"discount_price" db:"discount_price,omitempty" param:"discount_price"`
	Stock         int64     `json:"stock" db:"stock,omitempty" param:"stock"`
	ImageURL      string    `json:"image_url" db:"image_url,omitempty" param:"image_url"`
	IsDeleted     int64     `db:"is_deleted" json:"is_deleted,omitempty" param:"is_deleted"`
	CreatedAt     time.Time `db:"created_at" json:"created_at,omitempty" param:"created_at"`
	CreatedBy     string    `db:"created_by" json:"created_by,omitempty" param:"created_by"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at,omitempty" param:"updated_at"`
	UpdatedBy     string    `db:"updated_by" json:"updated_by,omitempty" param:"updated_by"`
	DeletedAt     time.Time `db:"deleted_at" json:"deleted_at,omitempty" param:"deleted_at"`
	DeletedBy     string    `db:"deleted_by" json:"deleted_by,omitempty" param:"deleted_by"`
}
