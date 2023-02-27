package routes

import (
	"net/http"

	"main.go/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
