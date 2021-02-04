package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/onunez-g/gorecipeapi/config"
	"github.com/onunez-g/gorecipeapi/data"
	"github.com/onunez-g/gorecipeapi/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Failed to load env vars")
	}
	cfg := config.Get()

	data.ConnectDatabase(cfg.GetDBConnStr())
	defer data.CloseConnection()

	r := router.Get()

	data.AutoMigrate()

	log.Println("Server listening...")
	apiPort := cfg.GetAPIPort()
	err := http.ListenAndServe(apiPort, r)
	if err != nil {
		panic(err.Error())
	}

}
