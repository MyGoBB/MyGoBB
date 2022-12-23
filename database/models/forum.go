package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Forum struct {
	Id          uint64 `gorm:"primaryKey"`
	ParentId    uint64
	Name        string
	Description sql.NullString
	IsLink      bool `gorm:"default:false"`
	Link        sql.NullString
	Forums      []Forum `gorm:"foreignkey:ParentId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
