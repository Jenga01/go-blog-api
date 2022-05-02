package server

import (
	"first/Config"
	"first/Routes"
	"first/Seeder"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var err error

func Run() {

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	Config.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Println("Opening status:", err)
	}

	Seeder.Load(Config.DB)

	r := Routes.InitializeRoutes()
	_ = r.Run()
}
