package repository

import (
	"github.com/tehdev/summoner-rift-api/entities"
	_playerCoinModel  "github.com/tehdev/summoner-rift-api/pkg/playerCoin/model"
)



type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string)(*_playerCoinModel.PlayerCoinShowing,error)
	

}