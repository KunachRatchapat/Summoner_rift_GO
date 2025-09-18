package repository

import "github.com/tehdev/summoner-rift-api/entities"

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	

}