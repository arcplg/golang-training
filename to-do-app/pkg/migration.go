package db

import (
	"database/sql"
	"log"
)

// MigrateToDoListTable creates the `to_do_list` table if it doesn't exist
func MigrateToDoListTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS to_do_list (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		is_completed BOOLEAN DEFAULT FALSE,
		user_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error migrating table: %v", err)
	}

	log.Println("Table `to_do_list` migration completed successfully!")
}

func MigrateUserTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error migrating table: %v", err)
	}

	log.Println("Table `users` migration completed successfully!")
}
