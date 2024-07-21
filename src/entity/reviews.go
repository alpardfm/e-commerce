package entity

import "time"

type Reviews struct {
	ID        int64     `json:"id" db:"id,omitempty" param:"id"`
	UserID    int64     `json:"user_id" db:"user_id" param:"user_id"`
	ProductID int64     `json:"product_id" db:"product_id,omitempty" param:"product_id"`
	Rating    int64     `json:"rating" db:"rating,omitempty" param:"rating"`
	Comment   string    `json:"comment" db:"comment,omitempty" param:"comment"`
	IsDeleted int64     `db:"is_deleted" json:"is_deleted,omitempty" param:"is_deleted"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty" param:"created_at"`
	CreatedBy string    `db:"created_by" json:"created_by,omitempty" param:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty" param:"updated_at"`
	UpdatedBy string    `db:"updated_by" json:"updated_by,omitempty" param:"updated_by"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at,omitempty" param:"deleted_at"`
	DeletedBy string    `db:"deleted_by" json:"deleted_by,omitempty" param:"deleted_by"`
}
