package main

import (
    "log"

    "github.com/go-sql-driver/mysql"

    "github.com/hd2yao/ecom/cmd/api"
    "github.com/hd2yao/ecom/config"
    "github.com/hd2yao/ecom/db"
)

func main() {
    cfg := mysql.Config{
        User:                 config.Envs.DBUser,
        Passwd:               config.Envs.DBPassword,
        Addr:                 config.Envs.DBAddress,
        DBName:               config.Envs.DBName,
        Net:                  "tcp",
        AllowNativePasswords: true,
        ParseTime:            true,
    }
    db, err := db.NewMySQLStorage(cfg)
    if err != nil {
        log.Fatal(err)
    }

    server := api.NewAPIServer(":8080", db)
    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}
