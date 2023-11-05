package models

type ManagementInvestor struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(300)" json:"name"`
	Position string `gorm:"type:varchar(300)" json:"position"`
	Image    string `gorm:"type:varchar(300)" json:"image"`
	Sequence int64  `json:"sequence"`
	Status   bool   `json:"status"`
}
