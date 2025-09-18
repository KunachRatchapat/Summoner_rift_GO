package main

import (
	"github.com/tehdev/summoner-rift-api/config"
	"github.com/tehdev/summoner-rift-api/databases"
	"github.com/tehdev/summoner-rift-api/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	//Run func
	playerMigration(tx)
	adminMigration(tx)
	inventoryMigration(tx)
	playercoinMigration(tx)
	cardMigration(tx)
	purcheaseHistoryMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func playerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Player{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func playercoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PlayerCoin{})
}

func cardMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Card{})
}

func purcheaseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
