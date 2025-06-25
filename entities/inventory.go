package entities

import "time"

type (
	Inventory struct {
		ID        uint64    	`gorm:"primaryKey;autoIncrement;"`
		Player_ID string   	 	`gorm:"type:varchar(64);not null;"`
		Card_ID   string   	 	`gorm:"type:varchar(128);not null;"`
		DeletedAt bool      	`gorm:"not null;default:false"`
		CreateAt  time.Time 	`gorm:"type:not null; autoCreateTime;"`
	}
)
