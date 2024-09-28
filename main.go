package main

import (
	"log"

	"github.com/rkpundhir90/task-management/api"
	"github.com/rkpundhir90/task-management/config"
	"github.com/rkpundhir90/task-management/db"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the database
	database, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.CloseDB(database)

	// Setup the router and run the server
	router := api.NewRouter(database)
	log.Println("Server is running on port", cfg.ServerPort)
	log.Fatal(router.Run(":" + cfg.ServerPort))
}
