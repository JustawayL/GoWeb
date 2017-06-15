package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

//Variable utilizada para cargar plantillas
var tpl *template.Template

//funcion encargada de cargar las plantillas ubicadas en la carpeta "templates"
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/acerca", acerca)
	http.HandleFunc("/subir", subir)
	http.HandleFunc("/subirProceso", subirProceso)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

//funcion encargada de responder la peticion a la ruta /index
func index(w http.ResponseWriter, r *http.Request) {
	searchDir := "public/pics"

	fileList := []string{}
	filenames := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range fileList {
		_, file := filepath.Split(p)
		filenames = append(filenames, file)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", filenames[1:])

}

//funcion encargada de responder la peticion a la ruta /acerca
func acerca(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "acerca.gohtml", nil)
}

//funcion encargada de responder la peticion a la ruta /subir
func subir(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "subir.gohtml", nil)
}

//funcion encargada de responder la peticion a la ruta /subirProceso
//Acepta el formulario POST para subir una nueva imagen
func subirProceso(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {

		t, _ := template.ParseFiles("subir.gohtml")
		t.Execute(w, nil)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		f, err := os.OpenFile("public/pics/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

	tpl.ExecuteTemplate(w, "subir.gohtml", "Imagen subida con exito")

}
