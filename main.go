package main

import (
	"log"

	"github.com/rkpundhir90/task-management/task-management/api"
	"github.com/rkpundhir90/task-management/task-management/app"
	"github.com/rkpundhir90/task-management/task-management/config"
	"github.com/rkpundhir90/task-management/task-management/db"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.CloseDB(database)

	appInstance := app.Build(database, &cfg)

	router := api.NewRouter(appInstance, &cfg)
	log.Println("Server is running on port", cfg.ServerPort)
	log.Fatal(router.Run(":" + cfg.ServerPort))
}
