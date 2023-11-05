package models

type SpecialOffer struct {
	Id     int64  `gorm:"primaryKey" json:"id"`
	Image  string `gorm:"type:varchar(300)" json:"image"`
	Url    string `gorm:"type:varchar(300)" json:"url"`
	Status bool   `json:"status"`
}
