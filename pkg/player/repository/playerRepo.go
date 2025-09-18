package repository

import "github.com/tehdev/summoner-rift-api/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error) 
	FindByID(playerID string) (*entities.Player, error)
}