package Database

import (
	"errors"
	"log"
	"os"

	"github.com/Tholkappiar/go-crud.git/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("err loading: %v", err)
    }
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		errors.New("missing addres")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	  }

	  DB = db
	  err = DB.AutoMigrate(&model.Blog{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected successfully!")
}