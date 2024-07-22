package entity

import "time"

type Users struct {
	ID        int64     `db:"id" json:"id,omitempty" param:"id"`
	Username  string    `db:"username" json:"username,omitempty" param:"username"`
	Email     string    `db:"email" json:"email,omitempty" param:"email"`
	Password  string    `db:"password" json:"password,omitempty" param:"password"`
	Pincode   string    `db:"pincode" json:"pincode,omitempty" param:"pincode"`
	RoleID    int64     `db:"role_id" json:"role_id,omitempty" param:"role_id"`
	IsActive  int64     `db:"is_active" json:"is_active,omitempty" param:"is_active"`
	IsDeleted int64     `db:"is_deleted" json:"is_deleted,omitempty" param:"is_deleted"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty" param:"created_at"`
	CreatedBy string    `db:"created_by" json:"created_by,omitempty" param:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty" param:"updated_at"`
	UpdatedBy string    `db:"updated_by" json:"updated_by,omitempty" param:"updated_by"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at,omitempty" param:"deleted_at"`
	DeletedBy string    `db:"deleted_by" json:"deleted_by,omitempty" param:"deleted_by"`
}
