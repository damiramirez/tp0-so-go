package main

import (
	"log"
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)


	log.Println("Running server on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}
