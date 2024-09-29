package db

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rkpundhir90/task-management/task-management/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database using GORM
func InitDB(cfg config.EnvironmentConfiguration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected!")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	schemaFile, err := os.Open("./task-management/db/schema.sql")
	if err != nil {
		log.Fatal("Failed to read schema file:", err)
	}
	defer schemaFile.Close()

	var schemaContent []byte
	buf := make([]byte, 1024)
	for {
		n, err := schemaFile.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal("Failed to read schema file:", err)
		}
		if n == 0 {
			break
		}
		schemaContent = append(schemaContent, buf[:n]...)
	}

	_, err = sqlDB.Exec(string(schemaContent))
	if err != nil {
		log.Fatal("Failed to execute schema:", err)
	}

	fmt.Println("Schema and sample data created successfully")

	// Retrieve and print sample data (optional)
	rows, err := sqlDB.Query(`SELECT t.id, t.title, t.description, u.name as assigned_user, t.status, t.due_date 
						   FROM tasks t 
						   LEFT JOIN users u ON t.user_id = u.id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Sample tasks:")
	for rows.Next() {
		var id int
		var title, description, user, status string
		var dueDate string

		err = rows.Scan(&id, &title, &description, &user, &status, &dueDate)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Task ID: %d\nTitle: %s\nDescription: %s\nAssigned User: %s\nStatus: %s\nDue Date: %s\n\n",
			id, title, description, user, status, dueDate)
	}

	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
