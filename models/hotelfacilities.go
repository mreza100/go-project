package models

type HotelFacilities struct {
	HotelFacilitiesId int64  `gorm:"primaryKey" json:"hotel_facilities_id"`
	HotelId           int64  `gorm:"type:varchar(300)" json:"hotel_id"`
	FacilityName      string `gorm:"type:varchar(300)" json:"facility_name"`
	Title             string `gorm:"type:varchar(300)" json:"title"`
	Image             string `gorm:"type:varchar(300)" json:"image"`
	Icon              string `gorm:"type:varchar(300)" json:"icon"`
	Sequence          int64  `json:"sequence"`
	Status            bool   `json:"status"`
}
