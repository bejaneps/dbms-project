package auth

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// a path to database of 'university.db'
var databasePath string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	databasePath = wd + "/db/university.db"
}

// GetDB returns a connection to sqlite database
func GetDB() *gorm.DB {
	DB, err := gorm.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return DB
}
