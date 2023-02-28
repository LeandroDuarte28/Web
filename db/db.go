package db

import (
	"database/sql"

	_ "github.com/lib/pq" // importa o driver PostgreSQL
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=db_Banco password=123456 sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}
