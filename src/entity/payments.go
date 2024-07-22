package entity

import "time"

type Payments struct {
	ID            int64     `db:"id" json:"id,omitempty" param:"id"`
	OrderID       int64     `db:"order_id" json:"order_id,omitempty" param:"order_id"`
	PaymentMethod string    `db:"payment_method" json:"payment_method,omitempty" param:"payment_method"`
	PaymentStatus string    `db:"payment_status" json:"payment_status,omitempty" param:"payment_status"`
	TransactionID string    `db:"transaction_id" json:"transaction_id,omitempty" param:"transaction_id"`
	IsDeleted     int64     `db:"is_deleted" json:"is_deleted,omitempty" param:"is_deleted"`
	CreatedAt     time.Time `db:"created_at" json:"created_at,omitempty" param:"created_at"`
	CreatedBy     string    `db:"created_by" json:"created_by,omitempty" param:"created_by"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at,omitempty" param:"updated_at"`
	UpdatedBy     string    `db:"updated_by" json:"updated_by,omitempty" param:"updated_by"`
	DeletedAt     time.Time `db:"deleted_at" json:"deleted_at,omitempty" param:"deleted_at"`
	DeletedBy     string    `db:"deleted_by" json:"deleted_by,omitempty" param:"deleted_by"`
}
