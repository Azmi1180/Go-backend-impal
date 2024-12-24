package main

import (
	"Intern_Backend/config"
	"Intern_Backend/docs"
	"Intern_Backend/routes"
	"Intern_Backend/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	// Load environment variables
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Programmatically set Swagger info
	docs.SwaggerInfo.Title = "Test API"
	docs.SwaggerInfo.Description = "Testing API Produk."
	docs.SwaggerInfo.Version = "1.0"

	// Set host from environment variable
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		log.Fatal("SWAGGER_HOST environment variable not set")
	}
	docs.SwaggerInfo.Host = swaggerHost

	// Specify supported schemes
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Connect to the database
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Setup and run the router
	r := routes.SetupRouter(db)
	r.Run(":80")
}
