package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	// "gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	// "github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	cnn, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/wallet_db")
	if err != nil {
		log.Fatal(err)
	}

	id := 1
	var wallet string

	if err := cnn.QueryRow("SELECT wallet FROM wallets WHERE id = ? LIMIT 1", id).Scan(&wallet); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, wallet)
}
