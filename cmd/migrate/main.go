package main

import (
    "log"
    "os"

    mysqlDriver "github.com/go-sql-driver/mysql"
    "github.com/golang-migrate/migrate/v4"
    mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"

    "github.com/hd2yao/ecom/config"
    "github.com/hd2yao/ecom/db"
)

func main() {
    cfg := mysqlDriver.Config{
        User:                 config.Envs.DBUser,
        Passwd:               config.Envs.DBPassword,
        Addr:                 config.Envs.DBAddress,
        DBName:               config.Envs.DBName,
        Net:                  "tcp",
        AllowNativePasswords: true,
        ParseTime:            true,
    }

    // 创建数据库连接实例
    db, err := db.NewMySQLStorage(cfg)
    if err != nil {
        log.Fatal(err)
    }

    // 创建数据库驱动实例
    driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // 创建迁移实例
    m, err := migrate.NewWithDatabaseInstance(
        "file://cmd/migration/migrations", // 指向迁移文件夹的 URL
        "mysql",                           // 数据库类型
        driver,
    )
    if err != nil {
        log.Fatal(err)
    }

    cmd := os.Args[len(os.Args)-1]
    if cmd == "up" {
        if err := m.Up(); err != nil && err != migrate.ErrNoChange {
            log.Fatal(err)
        }
    }
    if cmd == "down" {
        if err := m.Down(); err != nil && err != migrate.ErrNoChange {
            log.Fatal(err)
        }
    }
}
