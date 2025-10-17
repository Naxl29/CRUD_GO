package main

import (
    "net/http"
    "log"
    //"fmt"
	"text/template"

    //"CRUD_GO/config"
    //"CRUD_GO/routes"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Página Inicio")
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
 