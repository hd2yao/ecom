package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hd2yao/ecom/service/product"
	"github.com/hd2yao/ecom/service/user"
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
    subRouter := router.PathPrefix("/api/v1").Subrouter()

    // 注册 user 功能的路由以及处理函数
    userStore := user.NewStore(a.db)
    userHandler := user.NewHandler(userStore)
    userHandler.RegisterRoutes(subRouter)

    // 注册 product 功能的路由以及处理函数
    productStore := product.NewStore(a.db)
    productHandler := product.NewHandler(productStore)
    productHandler.RegisterRoutes(subRouter)

    log.Println("Listening on", a.addr)
    // 监听服务器 并为路由提供解决方法
    return http.ListenAndServe(a.addr, router)
}
