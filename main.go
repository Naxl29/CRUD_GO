package main

import (
    "net/http"
    "log"
    "fmt"
	"text/template"

    "CRUD_GO/config"
    "CRUD_GO/routes"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Inicio)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}
func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Página Inicio")
	tmpl.ExecuteTemplate(w, "Inicio", nil)
}
 