package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/tehdev/summoner-rift-api/databases"
	"github.com/tehdev/summoner-rift-api/entities"
	_playerException "github.com/tehdev/summoner-rift-api/pkg/player/exception"
)

type playerRepositorympl struct {
	db     databases.Database
	logger echo.Logger
}

func NewplayerRepositorympl(
	db databases.Database,
	logger echo.Logger,
) PlayerRepository {
	return &playerRepositorympl{
		db:     db,
		logger: logger,
	}
}


func (r *playerRepositorympl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	players := new(entities.Player)

	if err := r.db.Connect().Create(playerEntity).Scan(players).Error; err != nil{
		r.logger.Errorf("Creating Player failed: %s", err.Error())
		return  nil, &_playerException.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return  players, nil
  
}

func (r *playerRepositorympl) FindByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Where("id = ?", playerID).First(player).Error; err != nil{
		r.logger.Errorf("Find player by ID failed: %s",err.Error())
		return nil, &_playerException.PlayerNotFound{PlayerID: playerID}
	}

	return  player, nil
    
}