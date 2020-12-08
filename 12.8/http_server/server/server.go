package server


import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {

	 fmt.Fprintf(w, "hello\n")
}

func Sum(w http.ResponseWriter, req *http.Request) {

	 fmt.Fprintf(w, "sum\n")
}

func Headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func RunHttp(){

	http.ListenAndServe(":8080", nil)
}

