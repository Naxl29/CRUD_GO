package routes

import (
    "fmt"
    "net/http"
)

func SetupRoutes() *http.ServeMux {
    r := http.NewServeMux()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hola desde Go")
    })
    return r
}
