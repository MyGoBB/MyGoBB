package models

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	ForumId   uint64
	Forum     *Forum `gorm:"foreignKey:ForumId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
