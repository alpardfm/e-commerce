package entity

import "time"

type Products struct {
	ID            int64     `db:"id" json:"id,omitempty" param:"id"`
	CategoryID    int64     `db:"category_id" json:"category_id,omitempty" param:"category_id"`
	Name          string    `db:"name" json:"name,omitempty" param:"name"`
	Description   string    `db:"description" json:"description,omitempty" param:"description"`
	Price         float64   `db:"price" json:"price,omitempty" param:"price"`
	DiscountPrice float64   `db:"discount_price" json:"discount_price,omitempty" param:"discount_price"`
	Stock         int64     `db:"stock" json:"stock,omitempty" param:"stock"`
	ImageURL      string    `db:"image_url" json:"image_url,omitempty" param:"image_url"`
	IsDeleted     int64     `db:"is_deleted" json:"is_deleted,omitempty" param:"is_deleted"`
	CreatedAt     time.Time `db:"created_at" json:"created_at,omitempty" param:"created_at"`
	CreatedBy     string    `db:"created_by" json:"created_by,omitempty" param:"created_by"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at,omitempty" param:"updated_at"`
	UpdatedBy     string    `db:"updated_by" json:"updated_by,omitempty" param:"updated_by"`
	DeletedAt     time.Time `db:"deleted_at" json:"deleted_at,omitempty" param:"deleted_at"`
	DeletedBy     string    `db:"deleted_by" json:"deleted_by,omitempty" param:"deleted_by"`
}
