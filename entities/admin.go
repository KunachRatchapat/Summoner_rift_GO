package entities

import "time"

type Admin struct {
		ID       uint64 	`gorm:"primaryKey;autoIncrement;"`
		Cards 	 []Card  	`gorm:"foreignKey:AdminID;references:ID;constriant:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Name     string	 	`gorm:"type:varchar(128);not null;"`
		Email    string 	`gorm:"type:varchar(128);unique;not null"`
		Avatar   string 	`gorm:"type:varchar(256);not null;default:'';"`
		CreateAt time.Time 	`gorm:"not noull;autoCreateTime;"`
}
	
