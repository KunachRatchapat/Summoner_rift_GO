package entities

import "time"

type (
	Admin struct {
		ID       uint64 	`gorm:"primaryKey;autoIncrement;"`
		Name     string	 	`gorm:"type:varchar(128);"`
		Email    string 	`gorm:"type:varchar(128);"`
		Avatar   string 	`gorm:"type:varchar(256);"`
		CreateAt time.Time 	`gorm:"not noull;autoCreateTime;"`
	}
)
