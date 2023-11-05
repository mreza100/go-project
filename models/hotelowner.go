package models

type HotelOwner struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	HotelId string `gorm:"varchar(300)" json:"hotel_id" `
	UserId  int64  `json:"user_id"`
}
