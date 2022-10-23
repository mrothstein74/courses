package main

import (
	"fmt"

	_ "github.com/hellofresh/health-go/v4/checks/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/jackc/pgx/v4/stdlib"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "baloo"
	password = "junglebook"
	dbname   = "lenslocked"
)

type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;uniqueIndex"`
	Color  string
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	//db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{}, &Order{})

	var u User
	if err := db.Preload("Orders").First(&u).Error; err != nil {
		panic(err)
	}
	//createOrder(db, u, 1001, "fake description #1")
	//createOrder(db, u, 9999, "fake description #2")
	//createOrder(db, u, 100, "fake description #3")

}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error
	if err != nil {
		panic(err)
	}

}
