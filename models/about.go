package models

type About struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(300)" json:"title"`
	Image       string `gorm:"type:varchar(300)" json:"image"`
	BodyContent string `gorm:"varchar(300)" json:"body_content"`
}
