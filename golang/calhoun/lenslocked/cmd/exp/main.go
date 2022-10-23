package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	Name  string
	Email string `gorm:"not null;uniqueIndex"`
	Color string
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
	db.AutoMigrate(&User{})

	name, email, color := getInfo()
	u := User{
		Name:  name,
		Email: email,
		Color: color,
	}
	if err = db.Create(&u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)
}

func getInfo() (name, email, color string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("What is your email address?")
	email, _ = reader.ReadString('\n')
	fmt.Println("What is your favorite color?")
	color, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	color = strings.TrimSpace(color)
	return name, email, color
}
