package repository

import (
	"github.com/labstack/echo/v4"
	database "github.com/tehdev/summoner-rift-api/databases"
	"github.com/tehdev/summoner-rift-api/entities"
	_playerCoinException"github.com/tehdev/summoner-rift-api/pkg/playerCoin/exception"
	_playerCoinModel  "github.com/tehdev/summoner-rift-api/pkg/playerCoin/model"
)

type playerCoinRepositorympl struct{
	db database.Database
	logger echo.Logger
}

func NewPlayerCoinRepositorympl (db database.Database, logeer echo.Logger) PlayerCoinRepository{
	return &playerCoinRepositorympl{
		db: db,
		logger: logeer,
	}

}

func (r *playerCoinRepositorympl) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error){
	playerCoin := new(entities.PlayerCoin)

	if err := r.db.Connect().Create(playerCoinEntity).Scan(playerCoin).Error; err != nil{
		return nil, &_playerCoinException.AddingCoin{}
	}
	return playerCoin,nil
}

func (r *playerCoinRepositorympl) Showing(playerID string)(*_playerCoinModel.PlayerCoinShowing, error){
	playerCoinShowing:= new(_playerCoinModel.PlayerCoinShowing)

	if err := r.db.Connect().Model(&entities.PlayerCoin{}).Where("player_id = ?",playerID,).Select("player_id, sum(amount) as coin",
	).Group("player_id",).Scan(playerCoinShowing).Error; err != nil{
		r.logger.Errorf("player coin Showing failed: %s ", err.Error())
		return nil,&_playerCoinException.PlayerCoinShow{}
	}

	return  playerCoinShowing,nil
}

	