package main

import (
	"net/http"

	"main.go/routes"
	//"text/template"
	//"main.go/db"
	//"main.go/models"
)

func main() {
	//db := db.ConectaComBancoDeDados()
	//defer db.Close()
	//http.HandleFunc("/", index)
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
