package entities

import "time"

type (
	Inventory struct {
		ID        	uint64    		`gorm:"primaryKey;autoIncrement;"`
		PlayerID 	string   	 	`gorm:"type:varchar(64);not null;"`
		CardID   	string   	 	`gorm:"type:bigint;not null;"`
		IsDeleted 	bool      		`gorm:"not null;default:false;"`
		CreateAt  	time.Time 		`gorm:"not null; autoCreateTime;"`
	}
)
