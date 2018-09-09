package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	id           int
	nomeUsuario  string
	loginUsuario string
}

func main() {
	db, err := sql.Open("mysql", "masilva:spfc251191@(mysql472.umbler.com:41890)/barbearialucas")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	rows, err := db.Query("SELECT idUsuario, nomeUsuario, loginUsuario FROM usuario")

	for rows.Next() {
		var tag Tag
		err = rows.Scan(&tag.id, &tag.nomeUsuario, &tag.loginUsuario)
		if err != nil {
			panic("deu erro" + err.Error())
		}
		fmt.Println(tag.id, tag.nomeUsuario, tag.loginUsuario)
	}
}
