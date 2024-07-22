package entity

import "time"

type Cart struct {
	ID        int64     `db:"id" json:"id,omitempty" param:"id"`
	UserID    int64     `db:"user_id" json:"user,omitempty" param:"user_id"`
	ProductID int64     `db:"product_id" json:"product_id,omitempty" param:"product_id"`
	Quantity  int64     `db:"quantity" json:"quantity,omitempty" param:"quantity"`
	IsDeleted int64     `db:"is_deleted" json:"is_deleted,omitempty" param:"is_deleted"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty" param:"created_at"`
	CreatedBy string    `db:"created_by" json:"created_by,omitempty" param:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty" param:"updated_at"`
	UpdatedBy string    `db:"updated_by" json:"updated_by,omitempty" param:"updated_by"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at,omitempty" param:"deleted_at"`
	DeletedBy string    `db:"deleted_by" json:"deleted_by,omitempty" param:"deleted_by"`
}
