package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

///// Global variables /////
// initialize logs
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type DBConnection struct {
	db_conn *sql.DB
	err     string `none`
}

// Struct to store credrentials for use with db
type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

type UserCredentials struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Wallet      string `json:"wallet"`
	Transaction int    `json:"transaction"`
	Status      string
	DB_conn     *sql.DB
}

var (
	ctx context.Context
	db  *sql.DB
)

///// Global variables /////

///// Functions area begins /////
// FromJSON to be used for unmarshalling of user cred
func FromJSON(data []byte) UserCredentials {
	user_cred := UserCredentials{}
	err := json.Unmarshal(data, &user_cred)
	if err != nil {
		panic(err)
	}
	return user_cred
}

// Currently not used
func ExampleDB_PingContext() {
	// Ping and PingContext may be used to determine if communication with
	// the database server is still possible.
	//
	// When used in a command line application Ping may be used to establish
	// that further queries are possible; that the provided DSN is valid.
	//
	// When used in long running service Ping may be part of the health
	// checking system.

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	status := "up"
	if err := db.PingContext(ctx); err != nil {
		status = "down"
	}
	log.Println(status)
}

// Initialize log files
func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Enstablishing query connection
func (d *DBConnection) InnitDBConn() {
	tempdb, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/wallet_db")
	if err != nil {
		InfoLogger.Fatal(err)
	}
	d.db_conn = tempdb
	id := 1
	var wallet string

	// Simple query to check the connection
	if err := tempdb.QueryRow("SELECT wallet FROM wallets WHERE id = ? LIMIT 1", id).Scan(&wallet); err != nil {
		InfoLogger.Fatal(err)
	}
	fmt.Println(id, wallet)
	return
}

// Signup functions
func (d *DBConnection) Signup(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	InfoLogger.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err = d.db_conn.Query("insert into users values ($1, $2)", creds.Username, creds.Password); err != nil {
		InfoLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Query function
func RetrieveWallet(creds UserCredentials) int {
	var money int
	err := creds.DB_conn.QueryRow("SELECT money FROM wallets WHERE wallet=?", creds.Wallet).Scan(&money)
	if err != nil {
		InfoLogger.Fatal(err)
	}
	InfoLogger.Println("Wallet credits:", creds.Wallet)
	return money
}

// Update
func UpdateWallet(creds UserCredentials) int {
	var money int
	err := creds.DB_conn.QueryRow("UPDATE wallets SET money=? WHERE wallet=?", creds.Transaction, creds.Wallet).Scan(&money)
	if err != nil {
		InfoLogger.Fatal(err)
	}
	InfoLogger.Println("Wallet credits:", money)
	InfoLogger.Println(money)
	return money
}

// Handle wallet balance
func (b *UserCredentials) BalanceManage(w http.ResponseWriter, r *http.Request) {
	GetIP(r)
	// var final_balance int
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	err := json.NewDecoder(r.Body).Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	switch method := r.Method; method {
	case http.MethodPost:
		// b.Username = user_cred.Username
		// b.Password = user_cred.Password
		// b.Wallet = user_cred.Wallet
		// fmt.Fprint(w, b.Password)
		// fmt.Fprint(w, b.Username)
		// fmt.Printf(b.Wallet)
		// InfoLogger.Println("Requested wallet:", b.Wallet)
		// w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		current_balance := RetrieveWallet(*b)
		fmt.Println(current_balance - b.Transaction)
		// if current_balance-b.Transaction < 0 {
		// 	fmt.Println("Impossible transaction.")
		// } else {
		// 	b.Transaction = current_balance - b.Transaction
		// 	final_balance = UpdateWallet(*b) //has to return an int
		// 	fmt.Println(final_balance)
		// }
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

// get users IP
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	InfoLogger.Println("Requesters ip:", r.RemoteAddr)
	return r.RemoteAddr
}

// Main request and homepage request
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Homepage function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	InfoLogger.Println("Homepage requested")
	GetIP(r)
}

///// Functions area ends /////

// Main function
func main() {
	InfoLogger.Println("Starting the application...")
	d := &DBConnection{}
	d.InnitDBConn()

	b := &UserCredentials{}
	b.DB_conn = d.db_conn

	http.HandleFunc("/signup", d.Signup)
	http.HandleFunc("/manage-balance/input", b.BalanceManage)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// handleRequests()
}
