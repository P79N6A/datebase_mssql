package store

import (
    "flag"
    "fmt"
    "database/sql"
    _ "github.com/denisenkom/go-mssqldb"
    "log"
)



var debug = flag.Bool("debug", true, "enable debugging")
var password = flag.String("password", "Wanghh@345612", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "39.106.216.25", "the database server")
var user = flag.String("user", "SA", "the database user")
var database  = flag.String("database", "supermarket", "the database name")

var dB *sql.DB

func InitDB() {
    dB, err := GetDB()
    err = dB.Ping()
    if err != nil {
        log.Printf("connect db ping failed, err=%v", err)
    }
}

func GetDB() (*sql.DB, error) {
    if *debug {
        fmt.Printf(" password:%s\n", *password)
        fmt.Printf(" port:%d\n", *port)
        fmt.Printf(" server:%s\n", *server)
        fmt.Printf(" user:%s\n", *user)
    }
    connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable", *server, *database, *user, *password, *port)
     var err error
    db, err := sql.Open("mssql",connString)
    if err != nil {
        log.Printf("connect db failed, err=%v", err)
        return nil, err
    }
    return db, nil
}
