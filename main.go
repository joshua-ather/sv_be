package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/joshua-ather/sv_be/config"
	"github.com/joshua-ather/sv_be/database/migrations"
	"github.com/joshua-ather/sv_be/database/seeds"
	"github.com/joshua-ather/sv_be/routes"
	"log"
	"os"
)

var db *gorm.DB
var err error

func main() {

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	config.Database(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db = config.DB()

	route := routes.Init()

	handleArgs()

	route.Logger.Fatal(route.Start(":7444"))
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seeds.Execute(db, args[1:]...)
			return
		case "migrate":
			migrations.Execute(db, args[1:]...)
			return
		}
	}
}
