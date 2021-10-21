package database

import (
	"fmt"

	"github.com/ecam900/go-rest/internal/comment"
	"gorm.io/gorm"
)

// MigrateDB - migrates database and creates comment table
func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&comment.Comment{})
	if err != nil {
		fmt.Println("ERROR MIGRATING DATABASE")
		fmt.Println(err)
		return err
	}
	return nil
}
