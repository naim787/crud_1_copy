package db
import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  )
  
var DB *gorm.DB

func OPENDB(nmDB string) {
  var err error

	DB, err = gorm.Open(sqlite.Open(nmDB + ".db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&ListNote{},&ListMoneyBuy{})
}