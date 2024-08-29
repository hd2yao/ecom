package main

import (
    "log"

    "github.com/go-sql-driver/mysql"

    "github.com/hd2yao/ecom/cmd/api"
    "github.com/hd2yao/ecom/db"
)

func main() {
    cfg := mysql.Config{
        User:                 "root",
        Passwd:               "root",
        Addr:                 "127.0.0.1:3306",
        DBName:               "ecom",
        Net:                  "tcp",
        AllowNativePasswords: true,
        ParseTime:            true,
    }
    db, err := db.NewMySQLStorage(cfg)
    if err != nil {
        log.Fatal(err)
    }

    server := api.NewAPIServer(":8080", nil)
    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}
