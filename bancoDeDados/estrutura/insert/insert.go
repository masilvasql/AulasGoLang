package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO USUARIOS (nome) values (?)")
	stmt.Exec("Marcelo")
	stmt.Exec("Douglas")
	stmt.Exec("Gio")

	res, _ := stmt.Exec("Matheus")
	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
