package main

import (
    "database/sql"
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

    // 检查数据库是否连接成功
    initStorage(db)

    server := api.NewAPIServer(":8080", db)
    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}

func initStorage(db *sql.DB) {
    err := db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("DB: Successfully connected!")
}
