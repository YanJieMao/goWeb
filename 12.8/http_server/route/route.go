package route

import (
	"http_server/server"
	"net/http"
)

func InitRoute() {

	http.HandleFunc("/hello", server.Hello)
	http.HandleFunc("/headers", server.Headers)
	http.HandleFunc("/sum", server.Sum)

}
