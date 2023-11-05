package models

import "time"

type ArticleNews struct {
	Id                int64     `gorm:"primaryKey" json:"id"`
	ArticleCategoryId int64     `json:"article_category_id"`
	ImageThumbnail    string    `gorm:"varchar(300)" json:"image_thumbnail"`
	ImageBanner       string    `gorm:"varchar(300)" json:"image_banner"`
	Title             string    `gorm:"varchar(300)" json:"title"`
	Slug              string    `gorm:"varchar(300)" json:"slug"`
	MetaDescription   string    `gorm:"varchar(300)" json:"meta_description"`
	MetaKeyword       string    `gorm:"varchar(300)" json:"meta_keyword"`
	ContentNews       string    `gorm:"type:text" json:"content_news"`
	CreatedAt         time.Time `json:"created_at"`
	Status            bool      `json:"status"`
}
