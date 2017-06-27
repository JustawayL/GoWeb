package config

import(
"html/template"
)
//Variable utilizada para cargar plantillas
var Tpl *template.Template

//funcion encargada de cargar las plantillas ubicadas en la carpeta "plantillas"
func init() {
	Tpl = template.Must(template.ParseGlob("templates/*"))
}
