package entities

import "time"

type PurchaseHistory struct{
	ID  			uint64		`gorm:"primaryKey;autoIncrement;"`
	PlayerID		string		`gorm:"type:varchar(64);not null;"`
	CardID			uint64		`gorm:"type:bigint;not null;"`
	CardName		string 		`gorm:"type:varchar(64);not null;"`
	CardDescription	string		`gorm:"type:varchar(228);not null;"`
	CardPrice  		uint		`gorm:"not null;"`
	CardPicture		string		`gorm:"type:varchar(128);not null;"`
	Quantity		uint		`gorm:"not null;"`
	IsBuying		bool		`gorm:"type:boolean;not null;"`
	CreateAt		time.Time	`gorm:"not null;autoCreateTime;"`

}