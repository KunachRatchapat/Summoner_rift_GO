package entities

import(
	"time"
	_cardShopModel "github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
)

type (
	Card struct {
		ID 				uint64 			`gorm:"primaryKey;autoIncrement;"`
		AdminID			*string			`gorm:"type:varchar(64);"`	
		Name			string			`gorm:"type:varchar(64);unique;not null"`	
		Description		string			`gorm:"type:varchar(128);unique;not null;"`
		Picture			string			`gorm:"type:varchar(256);unique;not null;"`
		Price			int 			`gorm:"type:varchar(64);not null;"`
		DeleteAt		bool			`gorm:"not null;default:false;"`		
		CreateAt		time.Time		`gorm:"not null;autoCreateTime;"`
		UpdateAt		time.Time		`gorm:"not null;autoUpdateTime;"`
	}
)

func (i *Card) ToCardModel() *_cardShopModel.Card{
	return  &_cardShopModel.Card{
		ID: 	i.ID,
		Name:	i.Name,			
		Description: i.Description,	
		Picture: i.Picture,			
		Price: i.Price,	
	}
}
