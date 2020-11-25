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

func form(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form.html"),
	)
}

type Tarea struct {
	Nombre string
	Estado string
}

type AdminTareas struct {
	Tareas []Tarea
}

var misTareas AdminTareas

func (tareas *AdminTareas) Agregar(tarea Tarea) {
	tareas.Tareas = append(tareas.Tareas, tarea)
}

func (tareas *AdminTareas) String() string {
	var html string
	for _, tarea := range tareas.Tareas {
		html += "<tr>" +
			"<td>" + tarea.Nombre + "</td>" +
			"<td>" + tarea.Estado + "</td>" +
			"</tr>"
	}
	return html
}

func tareas(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		fmt.Println(req.PostForm)
		tarea := Tarea{Nombre: req.FormValue("tarea"), Estado: req.FormValue("estado")}
		misTareas.Agregar(tarea)
		fmt.Println(misTareas)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("respuesta.html"),
			tarea.Nombre,
		)
	case "GET":
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("tabla.html"),
			misTareas.String(),
		)
	}
}

func main() {
	http.HandleFunc("/form", form)
	http.HandleFunc("/tareas", tareas)
	fmt.Println("Corriendo servirdor de tareas...")
	http.ListenAndServe(":9000", nil)
}
