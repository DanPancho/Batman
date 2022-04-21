package main

import (
	"fmt"
	"gomysql/db"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Textoinicial string
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))

//Handler
func Index(res http.ResponseWriter, r *http.Request) {
	db.ConectarDb()
	inicio := Data{"> " + db.ObtenerAcertijo()}

	err := templates.ExecuteTemplate(res, "index.html", inicio)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(db.ObtenerAcertijo())
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
