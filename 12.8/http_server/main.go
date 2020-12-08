package main

import (
	"fmt"
	"http_server/route"
	"http_server/server"
)

func main() {

	route.InitRoute()
	fmt.Println("跑起来了")
	server.RunHttp()

}
