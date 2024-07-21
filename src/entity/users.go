package entity

import "time"

type Users struct {
	ID        int64     `json:"id" db:"id,omitempty" param:"id"`
	Username  string    `json:"username" db:"username,omitempty" param:"username"`
	Email     string    `json:"email" db:"email,omitempty" param:"email"`
	Password  string    `json:"password" db:"password,omitempty" param:"password"`
	Role      string    `json:"role" db:"role,omitempty" param:"role"`
	IsActive  int64     `json:"is_active" db:"is_active,omitempty" param:"is_active"`
	IsDeleted int64     `db:"is_deleted" json:"is_deleted,omitempty" param:"is_deleted"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty" param:"created_at"`
	CreatedBy string    `db:"created_by" json:"created_by,omitempty" param:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty" param:"updated_at"`
	UpdatedBy string    `db:"updated_by" json:"updated_by,omitempty" param:"updated_by"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at,omitempty" param:"deleted_at"`
	DeletedBy string    `db:"deleted_by" json:"deleted_by,omitempty" param:"deleted_by"`
}
