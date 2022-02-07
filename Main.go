package main

import (
	"net/http"
	"log"	
)
//Iniciando la aplicación y el puerto 8080
func main(){
	//MidleWare llamado de funciones de los enpoints
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

}