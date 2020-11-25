package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}

func root(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("index.html"),
	)
}

func main() {
	http.HandleFunc("/", root)
	fmt.Println("Arrancando el servidor...")
	http.ListenAndServe(":9000", nil)
}
