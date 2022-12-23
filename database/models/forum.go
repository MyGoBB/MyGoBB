package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Forum struct {
	Id          uint64 `gorm:"primaryKey"`
	ParentId    *uint64
	Name        string
	Description sql.NullString
	IsLink      bool `gorm:"default:false"`
	Link        sql.NullString
	Children    []Forum `gorm:"foreignkey:ParentId"`
	Topics      []Topic
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
