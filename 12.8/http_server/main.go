package main

import(
	"fmt"
	"http_server/server"
	"http_server/route"

)


func main() {

	route.InitRoute()
	fmt.Println("跑起来了")
	server.RunHttp()
	
}