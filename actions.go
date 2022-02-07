package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
)
//Peliculas disponible
var movies = Movies{
	Movie{"Sin Limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo Gas", 2005, "Juan Antonio"}, 
}
//Función master para listado o view de movies
func Index(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola Mundo desde mi servidor web con GO")
	}
	//Listando las peliculas
	func Movielist(w http.ResponseWriter, r *http.Request){

		//fmt.Fprintf(w, "Listado de Peliculas")
		//Conectando ala db para listar las peliculas
		json.NewEncoder(w).Encode(movies)
		
	}
	//Mostrando las peliculas
	func MovieShow(w http.ResponseWriter, r *http.Request){
		//Conectando a base de datos para ver las peliculas
		params := mux.Vars(r)
		movie_id := params["id"]

		fmt.Fprintf(w, "Has cargado la pelicula Numero %s", movie_id)
	}	

func MovieAdd(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if(err != nil){
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)	
}







