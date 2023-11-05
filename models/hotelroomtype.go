package models

type HotelRoomType struct {
	HotelRoomTypeId int64  `gorm:"primaryKey" json:"hotel_room_type_id"`
	HotelId         int64  `json:"hotel_id"`
	RoomName        string `gorm:"type:varchar(300)" json:"room_name"`
	Title           string `gorm:"type:varchar(300)" json:"title"`
	Image           string `gorm:"type:varchar(300)" json:"image"`
	Url             string `gorm:"type:varchar(300)" json:"url"`
	Sequence        int64  `json:"sequence"`
	Status          bool   `json:"status"`
}
