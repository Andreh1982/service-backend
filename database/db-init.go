package database

import (
	"service-backend/shared"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	host := "localhost"
	port := "5432"
	dbname := "root"
	user := "root"
	password := "root"
	db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" sslmode=disable password="+password)
	if err != nil {
		shared.LogCustom([]string{"Database Connection Failed!"}, "error")
	}
	db.LogMode(true)
	// db.AutoMigrate(models.Buyer{})
	// db.AutoMigrate(models.Product{})
	// db.AutoMigrate(models.Seller{})

	DB = db
	shared.LogCustom([]string{"Database Connected!"}, "info")
}

func GetDB() *gorm.DB {
	return DB
}

func ClearTable() {
	DB.Exec("DELETE FROM root")
}
