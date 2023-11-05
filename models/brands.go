package models

import (
	"time"

	"gorm.io/gorm"
)

type Brands struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	TitleBrand  string `gorm:"type:varchar(300)" json:"brand_title"`
	Description string `gorm:"text" json:"description"`
	Image       string `gorm:"varchar(300)" json:"image"`
	Url         string `gorm:"varchar(300)" json:"url"`
	Sequence    int64  `json:"sequence"`
	Status      bool   `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
