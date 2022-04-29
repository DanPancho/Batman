package main

import (
	"fmt"
	"gomysql/db"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//var outputs []db.Acertijo
// Variables Globales
var first bool = true
var acertijos []db.Acertijo
var intentos int = 0
var problema int = 1

type Front struct {
	Acertijo  string
	Respuesta string
}

var respuesta []Front

type Enviar struct {
	Datos []Front
}

// TEMPLATES
var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))

func cargarNuevoAcertijo(res http.ResponseWriter, mensaje string) {
	intentos += 1
	acertijos = append(acertijos, db.ObtenerAcertijo(problema)[0])
	respuesta = append(respuesta, Front{db.ObtenerAcertijo(problema)[0].Acertijo, mensaje})
	dato := Enviar{respuesta}
	err := templates.ExecuteTemplate(res, "index.html", dato)
	if err != nil {
		fmt.Print(err)
	}
}

//Handler
func Index(res http.ResponseWriter, r *http.Request) {

	if intentos <= 9 {
		if first {
			first = false
			cargarNuevoAcertijo(res, " ")

		} else {
			erro := r.ParseForm()
			if erro != nil {
				log.Panic(erro)
			}
			if input, ok := r.Form["input"]; ok {

				if strings.ToLower(input[0]) == acertijos[len(acertijos)-1].Respuesta {
					problema += 1
					//intentos += 1
					cargarNuevoAcertijo(res, "No pares...")
				} else {
					cargarNuevoAcertijo(res, "ERROR")
				}

				/*fmt.Println("INPUT -> ", strings.ToLower(input[0]))
				fmt.Println(acertijos)
				fmt.Println("BD -> ", acertijos[len(acertijos)-1].Respuesta)*/
			}
		}
	} else {
		const valor = 302
		http.Redirect(res, r, "https://www.rataalada.com/", valor)
	}

}

func main() {

	// Conexion
	db.ConectarDb()

	//Archivos estáticos
	archEstaticos := http.FileServer(http.Dir("estaticos"))

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	//Mux de archivos estáticos
	mux.Handle("/estaticos/", http.StripPrefix("/estaticos/", archEstaticos))

	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())

}
