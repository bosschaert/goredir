package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/foo", fooPage)
	http.ListenAndServe(":8888", nil)
}

func rootPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Serving root")
	request.ParseForm()
	writer.Header().Set("Location", "/foo")
	writer.WriteHeader(302)
}

func fooPage(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("Serving page: " + request.URL.Path)
	fmt.Fprintln(writer, `<h1>Foo!</h1>`)
}
