package api

import (
    "database/sql"
    "net/http"

    "github.com/gorilla/mux"
)

type APIServer struct {
    addr string
    db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
    return &APIServer{
        addr: addr,
        db:   db,
    }
}

func (a *APIServer) Run() error {
    // 初始化路由 注册所有路由
    // 在这里可以自定义路由及其处理函数
    // 只要实现了接口 Handler 就可以使用
    router := mux.NewRouter()

    // 监听服务器 并为路由提供解决方法
    return http.ListenAndServe(a.addr, router)
}
