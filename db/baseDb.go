package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:root@tcp(localhost:3306)/batman"

var conexionDb *sql.DB

type Acertijo struct {
	id        int
	Acertijo  string
	Respuesta string
}
type Acertijos []Acertijo

//abre la conexión a la base de datos
func ConectarDb() {
	cn, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión exitosa")
	conexionDb = cn
}

func ObtenerAcertijo(id int) []Acertijo {
	sql := "SELECT * FROM acertijos WHERE id = ?"
	acertijos := Acertijos{}
	if data, err := conexionDb.Query(sql, id); err != nil {
		fmt.Print(err)
	} else {
		for data.Next() {
			acertijo := Acertijo{}
			data.Scan(&acertijo.id, &acertijo.Acertijo, &acertijo.Respuesta)
			acertijos = append(acertijos, acertijo)
		}
	}
	return acertijos
}

//_ "github.com/go-sql-driver/mysql"
