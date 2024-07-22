package entity

import "time"

type Categories struct {
	ID        int64     `db:"id" json:"id,omitempty" param:"id"`
	Name      string    `db:"name" json:"name,omitempty" param:"name"`
	IsDeleted int64     `db:"is_deleted" json:"is_deleted,omitempty" param:"is_deleted"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty" param:"created_at"`
	CreatedBy string    `db:"created_by" json:"created_by,omitempty" param:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty" param:"updated_at"`
	UpdatedBy string    `db:"updated_by" json:"updated_by,omitempty" param:"updated_by"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at,omitempty" param:"deleted_at"`
	DeletedBy string    `db:"deleted_by" json:"deleted_by,omitempty" param:"deleted_by"`
}

type BodyCategories struct {
	Name string `json:"name"`
}

type PaginationCategories struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}

type ResponseCategories struct {
	Limit      int64        `json:"limit"`
	Page       int64        `json:"page"`
	TotalRows  int64        `json:"total_rows"`
	TotalPages int64        `json:"total_pages"`
	Data       []Categories `json:"data"`
}
