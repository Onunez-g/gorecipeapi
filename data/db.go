package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/onunez-g/gorecipeapi/models"
)

//Db database
var Db *gorm.DB
var err error

func ConnectDatabase(connStr string) {
	Db, err = gorm.Open("postgres", connStr)
	if err != nil {

		panic(err.Error())

	}
}

func CloseConnection() {
	Db.Close()
}

func AutoMigrate() {
	Db.AutoMigrate(&models.Ingredient{})
	Db.AutoMigrate(&models.RecipeDetail{})
	Db.AutoMigrate(&models.Recipe{})
}
