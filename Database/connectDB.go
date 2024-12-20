package Database

import (
	"fmt"
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
		fmt.Printf("No Database URL.")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	  }

	  DB = db
	  err = DB.AutoMigrate(&model.Blog{})
	if err != nil {
		fmt.Printf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database connected successfully!")
}