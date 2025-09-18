package entities

import(
	"time"
	_cardShopModel "github.com/tehdev/summoner-rift-api/pkg/cardShop/model"
)

type (
	
	Card struct {
    	ID          uint64    `gorm:"primaryKey;autoIncrement"`
    	AdminID     *string   `gorm:"type:varchar(64)"`
    	Name        string    `gorm:"type:varchar(64);unique;not null"`
    	Description string    `gorm:"type:varchar(128);not null"`
    	Picture     string    `gorm:"type:varchar(256);not null"`
    	Price       int       `gorm:"not null"`
    	IsArchive   bool      `gorm:"not null;default:false"` 
    	CreatedAt   time.Time `gorm:"autoCreateTime"`
    	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
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
