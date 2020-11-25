package main

import (
	"fmt"
	"net/http"
)

func root(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Ejemplo Hola</title>
			</head>
			<body>
				Hola Mundo!
			</body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/", root)
	fmt.Println("Arrancando el servidor...")
	http.ListenAndServe(":9000", nil)
}
