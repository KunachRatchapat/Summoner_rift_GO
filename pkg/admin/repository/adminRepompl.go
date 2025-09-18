package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/tehdev/summoner-rift-api/databases"
	"github.com/tehdev/summoner-rift-api/entities"
	_adminException "github.com/tehdev/summoner-rift-api/pkg/admin/exception"
)

type adminRepositorympl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositorympl(
	db databases.Database,
	logger echo.Logger,
) AdminRepository {
	return &adminRepositorympl{
		db:     db,
		logger: logger,
	}
}


func (r *adminRepositorympl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.Connect().Create(adminEntity).Scan(admin).Error; err != nil{
		r.logger.Errorf("Creating Player failed: %s", err.Error())
		return  nil, &_adminException.AdminCreating{AdminID: adminEntity.ID}
	}

	return admin, nil
}	

func (r *adminRepositorympl) FindByID(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.Connect().Where("id = ?", adminID).First(admin).Error; err != nil{
		r.logger.Errorf("Find admin by ID failed: %s",err.Error())
		return nil, &_adminException.AdminNotFound{AdminID: adminID}
	}

	return  admin, nil
    
}