package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jinzhu/gorm"
)

//Department struct represents a departments info in university
type Department struct {
	DeptName string `gorm:"type: TEXT; PRIMARY_KEY; NOT NULL"`
	Building string `gorm:"type: TEXT"`
	Budget   string `gorm:"type: TEXT"`
}

const databasePath = "/home/bezhan/Programming/go/src/github.com/bejaneps/dbms-project/db/university.db"
const csvPath = "/home/bezhan/Programming/go/src/github.com/bejaneps/dbms-project/utils/departments.csv"

// CopyFromCSV copies data from csv to db
func CopyFromCSV() {
	/* DATABASE CONNECTION */
	DB, err := gorm.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer DB.Close()

	/* CSV CONNECTION */
	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatalf("Unable to open a csv file: %v", err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		dep := Department{record[0], record[1], record[2]}
		DB.Table("department").Create(&dep)
	}
}
