package main

import (
	"net/http"
	"github.com/JustawayL/GoWeb/config"
)


func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", config.Index)
	http.HandleFunc("/acerca", config.Acerca)
	http.HandleFunc("/subir", config.Subir)
	http.HandleFunc("/subirProceso", config.SubirProceso)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
