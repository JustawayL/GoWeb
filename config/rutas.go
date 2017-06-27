package config

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

//funcion encargada de responder la peticion a la ruta /index
func Index(w http.ResponseWriter, r *http.Request) {
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

	Tpl.ExecuteTemplate(w, "index.gohtml", filenames[1:])

}

//funcion encargada de responder la peticion a la ruta /acerca
func Acerca(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "acerca.gohtml", nil)
}

//funcion encargada de responder la peticion a la ruta /subir
func Subir(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "subir.gohtml", nil)
}

//funcion encargada de responder la peticion a la ruta /subirProceso
//Acepta el formulario POST para subir una nueva imagen
func SubirProceso(w http.ResponseWriter, r *http.Request) {
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

	Tpl.ExecuteTemplate(w, "subir.gohtml", "Imagen subida con exito")

}
