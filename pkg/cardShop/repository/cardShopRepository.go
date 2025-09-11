package repository

import "github.com/tehdev/summoner-rift-api/entities"

type CardshopRepository interface {
	Listing() ([]*entities.Card, error)
}