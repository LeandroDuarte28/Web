package controllers

import (
	"net/http"
	"text/template"

	"main.go/models"
)

// aqui existe um erro pois não retornam os valores para o site
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}
