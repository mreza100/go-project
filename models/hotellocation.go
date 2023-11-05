package models

type HotelLocation struct {
	HotelLocationId int64  `gorm:"primaryKey" json:"hotel_location_id"`
	NameLocation    string `gorm:"type:varchar(300)" json:"name_location"`
	ImageDesktop    string `gorm:"type:varchar(300)" json:"image_desktop"`
	ImageMobile     string `gorm:"type:varchar(300)" json:"image_mobile"`
	Sequence        int64  `json:"sequence"`
	Status          bool   `json:"status"`
}
