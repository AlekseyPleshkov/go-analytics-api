package main

import (
	"go-analytics/internal/builder"
	"go-analytics/internal/config"
	"go-analytics/internal/data"
	"log"
)

func main() {
	config, err := config.NewConfig()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	db, err := data.NewConnect(config.DB)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	routes, err := builder.Create(config, db)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	routes.Run(config.Host)
}
