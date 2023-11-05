package models

type HotelInformation struct {
	HotelInformationId int64  `gorm:"primaryKey" json:"hotel_information_id"`
	HotelId            int64  `gorm:"type:varchar(300)" json:"hotel_id"`
	ShortDescription   string `gorm:"type:varchar(300)" json:"short_description"`
	Address            string `gorm:"type:varchar(300)" json:"address"`
	UrlGmaps           string `gorm:"type:varchar(300)" json:"url_gmaps"`
	Phone              string `gorm:"type:varchar(300)" json:"phone"`
	Email              string `gorm:"type:varchar(300)" json:"email"`
	Whatsapp           string `gorm:"type:varchar(300)" json:"whatsapp"`
	Instagram          string `gorm:"type:varchar(300)" json:"instagram"`
	Facebook           string `gorm:"type:varchar(300)" json:"facebook"`
	Sequence           int64  `json:"sequence"`
	Status             bool   `json:"status"`
}
