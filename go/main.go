package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
	// _ "github.com/go-sql-driver/mysql"
	// "gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	// "github.com/gin-gonic/gin"
	// "net/http"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	InfoLogger.Println("Homepage requested")
}

func main() {
	InfoLogger.Println("Starting the application...")
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

	handleRequests()
}
