package route

import (
	"net/http"
	"http_server/server"

)





func InitRoute() {

	http.HandleFunc("/hello", server.Hello)
	http.HandleFunc("/headers", server.Headers)
	http.HandleFunc("/sum", server.Sum)


}