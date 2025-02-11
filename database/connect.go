package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/LinhNguyen411/notes-api-fiber/config"
	"github.com/LinhNguyen411/notes-api-fiber/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("Idiot")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB_USER"),
		config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database")

	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database Migrate")
}
