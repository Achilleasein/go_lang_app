package main

import(
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
        "log"
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

        fmt.Println(id, name)
}