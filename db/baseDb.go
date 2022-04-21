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
	acertijo  string
	respuesta string
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

func ObtenerAcertijo() string {
	sql := "SELECT * FROM acertijos WHERE id = 2"
	acertijos := Acertijos{}
	if data, err := conexionDb.Query(sql); err != nil {
		fmt.Print(err)
	} else {
		for data.Next() {
			acertijo := Acertijo{}
			data.Scan(&acertijo.id, &acertijo.acertijo, &acertijo.respuesta)
			acertijos = append(acertijos, acertijo)
		}
	}
	return acertijos[0].acertijo
}

//_ "github.com/go-sql-driver/mysql"
