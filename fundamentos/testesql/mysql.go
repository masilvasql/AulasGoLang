package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	idUsuario int
	deslogin  string
	dessenha  string
}

func main() {
	fmt.Println("Go MySQL Tutorial")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dbphp7")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	rows, err := db.Query("SELECT idUsuario, deslogin, dessenha FROM tb_usuarios")

	for rows.Next() {
		var tag Tag
		err = rows.Scan(&tag.idUsuario, &tag.deslogin, &tag.dessenha)
		if err != nil {
			panic("deu erro" + err.Error())
		}
		fmt.Println(tag.idUsuario, tag.deslogin, tag.dessenha)
	}
}
