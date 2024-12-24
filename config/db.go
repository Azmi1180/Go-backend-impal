package config

import (
	"Intern_Backend/models"
	"Intern_Backend/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := utils.Getenv("DBUSER", "root")
	password := utils.Getenv("DBPASS", "")
	host := utils.Getenv("DBHOST", "127.0.0.1")
	port := utils.Getenv("DBPORT", "3306")
	database := utils.Getenv("DBNAME", "industrial")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	if err := db.AutoMigrate(&models.AdminModel{}, &models.ManagerModel{}, &models.BarangModel{}); err != nil {
		panic(fmt.Sprintf("Database migration failed: %v", err))
	}

	return db
}
