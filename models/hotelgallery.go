package models

type HotelGallery struct {
	HotelGalleryId int64  `gorm:"primaryKey" json:"hotel_gallery_id"`
	HotelId        int64  `gorm:"type:varchar(300)" json:"hotel_id"`
	GalleryName    string `gorm:"type:varchar(300)" json:"gallery_name"`
	Image          string `gorm:"type:varchar(300)" json:"image"`
	Sequence       int64  `json:"sequence"`
	Status         bool   `json:"status"`
}
