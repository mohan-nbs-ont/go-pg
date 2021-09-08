package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

func main() {
    term_list := os.Getenv("TERM_LIST")
    dt := os.Getenv("DT")

    if term_list == "" || dt == "" {
        fmt.Fprintln(os.Stderr, "Required environment variables: TERM_LIST, DT")
        os.Exit(1)
    }

    db, err := sql.Open("postgres", "postgresql://")
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }
    defer db.Close()

    if _, err := db.Exec(`SELECT loadtms($1,$2,'full')`, term_list, dt); err != nil {
        log.Fatal(err)
    }
}
