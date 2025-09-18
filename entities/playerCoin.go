package entities

import (
	"time"
	_playerCoinModel "github.com/tehdev/summoner-rift-api/pkg/playerCoin/model"
)

type (
	PlayerCoin struct {
	ID       	uint64 	`gorm:"primaryKey;autoIncrement;"`
	PlayerID 	string 	`gorm:"type:varchar(64);not null;"`
	Amount		int64  	`gorm:"not null;"`
	CreateAt	time.Time `gorm:"autoCreateTime;not null;"`
}

)

func (p *PlayerCoin) ToPlayerCoinModel() *_playerCoinModel.PlayerCoin {
	return &_playerCoinModel.PlayerCoin{
		ID:        p.ID,
		PlayerID:  p.PlayerID,
		Amount:    p.Amount,
		CreatedAt: p.CreateAt,
	}

}


