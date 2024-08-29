package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    PublicHost string
    Port       string
    DBUser     string
    DBPassword string
    DBAddress  string
    DBName     string
}

var Envs = initConfig()

func initConfig() Config {
    // 默认文件名为 '.env'
    // 它会将这些环境变量加载到当前进程的环境变量中
    // 使它们可以通过标准库中的 os.Getenv 或 os.LookupEnv 等函数访问
    godotenv.Load()

    return Config{
        PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
        Port:       getEnv("PORT", "8080"),
        DBUser:     getEnv("DB_USER", "root"),
        DBPassword: getEnv("DB_PASSWORD", "mypassword"),
        DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
        DBName:     getEnv("DB_NAME", "ecom"),
    }
}

func getEnv(key, fallBack string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallBack
}
