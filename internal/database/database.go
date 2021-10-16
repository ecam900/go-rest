package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase - returns a pointer to a database object
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up Database Connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("RETURNING DB HERE")
		return db, err
	}

	// sqldb, err := db.DB()
	// if err == nil {
	// 	fmt.Println("ERROR")
	// 	fmt.Println(err.Error())
	// }

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("PINGING HERE!")

		sqlDB.Ping()
		fmt.Println(sqlDB)
	}

	return db, nil
}
