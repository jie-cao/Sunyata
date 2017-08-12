package main

import (
	"log"
	"net/http"
	"sunyata/core/route"
)

func main() {
	//http.HandleFunc("/", route.DispatchRoute) //设置访问的路由

	router := route.NewRouter()
	RegisterRoutes(router)
	
	err := http.ListenAndServe(":9090", router) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
