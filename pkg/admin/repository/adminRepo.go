package repository

import(
	"github.com/tehdev/summoner-rift-api/entities"
)
type AdminRepository interface{
	Creating(adminEntity *entities.Admin) (*entities.Admin, error) 
	FindByID(adminID string) (*entities.Admin, error)
}