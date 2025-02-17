package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID        string `json:"id" gorm:"index:idx_id,sort:asc"`
	Name      string `json:"name" binding:"required" gorm:"index:idx_name,sort:asc"`
	Nickname  string `json:"nickname" binding:"required" gorm:"index:idx_nickname,sort:asc,unique"`
	Birth     string `json:"birth" binding:"required"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
