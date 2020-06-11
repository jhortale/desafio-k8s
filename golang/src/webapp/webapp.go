package main

import (
	"fmt"
	"net/http"
)

func greeting(msg string) string {
	return "<b>" + msg + "</b>"
}

func HttpServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", HttpServer)
	http.ListenAndServe(":8000", nil)
}
