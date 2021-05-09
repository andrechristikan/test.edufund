package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Loan struct {
	gorm.Model
	Number		uint16		`gorm:"index, unique"`
	Amount		uint32
	Tenor		uint8
	Status		string		`gorm:"index"`
	Borrower	string		`gorm:"index"`
	ApproveDate	time.Time
}

//Init ...
func Init() {
	dbinfo := fmt.Sprintf(
		"host=%s port=%s user=%s database=%s sslmode=disable password=%s", 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"), 
	)
	
	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Loan{})
	
}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}
