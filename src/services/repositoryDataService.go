package services

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
    "hacker-rank-lambda/src"
	_ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
	fmt.Println("Testing connection...")
	db, err := sql.Open("postgres", config.DB_CONNECTION)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertData(db *sql.DB, filepath string) error {
    file, err := os.Open(filepath)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        _, err := db.Exec(scanner.Text())
        if err != nil {
            fmt.Printf("Error inserting data: %v\n", err)
            continue
        }
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    fmt.Println("Data inserted successfully!")
    return nil
}
