package entities

import "time"

type PlayerCoin struct {
	ID       	uint64 `gorm:"primaryKey;autoIncrement;"`
	PlayerID 	string `gorm:"not null;type:varchar(64);"`
	Amount		int64  `gorm:"not null;"`
	CreateAt	time.Time `gorm:"autoCreateTime;not null;"`
}
