package manipulador

import "html/template"

//ModeloOla armazena os modelos de página que serão executados pelos manipuladores
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazena os modelos de página que serão executados pelos manipuladores
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
