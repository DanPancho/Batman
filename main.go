package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Textoinicial string
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))

//Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	inicio := Data{`TRACEROUTE RATAALADA.COM 
	POS-0-3-0-0-CR01.ARKAM.GOTHAMDATA.NET`}
	err := templates.ExecuteTemplate(rw, "index.html", inicio)
	if err != nil {
		fmt.Print(err)
	}
}

func main() {

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
