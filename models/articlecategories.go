package models

type ArticleCategories struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	CategoryName string `gorm:"type:varchar(300)" json:"category_name"`
	Sequence     int64  `json:"sequence"`
	Status       bool   `json:"status"`
}
