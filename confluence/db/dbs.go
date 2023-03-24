package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type TestPage struct {
	Id    string
	Title string
	Body  string
}

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=dev password=%s dbname=pages port=5432 sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PASS"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}

func GetPages() *[]TestPage {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&TestPage{})
	//db.Table("pages").Where("gender = ?", "male").Limit(100).Find(&persons)
	var pages []TestPage
	db.Find(&pages)
	return &pages
}
