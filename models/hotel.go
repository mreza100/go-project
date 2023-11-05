package models

type Hotel struct {
	HotelId         int64  `gorm:"primaryKey" json:"hotel_id"`
	HotelLocationId int64  `gorm:"type:varchar(300)" json:"hotel_location_id"`
	BrandId         int64  `gorm:"type:varchar(300)" json:"brand_id"`
	ImageDesktop    string `gorm:"type:varchar(300)" json:"image_desktop"`
	ImageMobile     string `gorm:"type:varchar(300)" json:"image_mobile"`
	TitleHotel      string `gorm:"type:varchar(300)" json:"title_hotel"`
	Slug            string `gorm:"type:varchar(300)" json:"slug"`
	Description     string `gorm:"varchar(300)" json:"description"`
	Status          bool   `json:"status"`
}
