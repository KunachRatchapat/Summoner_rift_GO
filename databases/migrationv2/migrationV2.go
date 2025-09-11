// เอาไว้ generate card เพื่อสุ่ม card แล้วใส่ใน ดาต้าเบส
package main

import (
	"github.com/tehdev/summoner-rift-api/entities"
	"gorm.io/gorm"
	"github.com/tehdev/summoner-rift-api/databases"
	"github.com/tehdev/summoner-rift-api/config"
)

func main() {
	conf := config.ConfigGetting()
	db := 	databases.NewPostgresDatabase(conf.Database)

	tx := db.ConnectionGetting().Begin()

	//Run func
	cardAdding(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func cardAdding(tx *gorm.DB) {
	cards := []entities.Card{
		{
			Name: "Katarina",
			Description: "Magic Damage and Assasin Role can killed ADC",
			Price: 300,
			Picture: "https://surl.li/rsbnib",
		},

		{
			Name: "Ezreal",
			Description: "Can auto attack Tank, Figter, Assasin and Mage role is ADC Carry ",
			Price: 250,
			Picture: "https://surl.lt/gqxjaw",
		},

		{
			Name: "Darius",
			Description: "Can Easy killing Mage, Adc and your Tank so Ezzz !! Role Figther !",
			Price: 300,
			Picture: "https://surl.li/zrmilb",
		},

		{
			Name: "Orianna",
			Description: " Your team so Ezy killed your Adc, Tank and Support Role Mage and can control vision your team EiEi !",
			Price: 300,
			Picture: "https://surl.li/mjdrfo",
		},

		{
			Name: "Lulu",
			Description: "knocking nearby enemies into the air and granting the ally a large amount of bonus health Role Support !",
			Price: 200,
			Picture: "https://surl.li/msldws",
		},




		
	}

	tx.CreateInBatches(cards, len(cards)) //เลือกใช้ตัวนี้เพราะผมป้องกันระบบล่ม ถ้าหายตัวนึงก็หายหมดเลยอะ
}