// package config

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func LoadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// }

// func ConnectDB() {
// 	var err error

// 	dsn := "host=" + os.Getenv("DB_HOST") +
// 		" user=" + os.Getenv("DB_USER") +
// 		" password=" + os.Getenv("DB_PASSWORD") +
// 		" dbname=" + os.Getenv("DB_NAME") +
// 		" port=" + os.Getenv("DB_PORT") +
// 		" sslmode=disable"

//		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//		if err != nil {
//			log.Printf("failed to connect to database, got error %v", err)
//			log.Fatal("Failed to connect to the database.")
//		}
//	}
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func ConnectDB() {
	var err error

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to database, got error %v", err)
		log.Fatal("Failed to connect to the database.")
	}
}
