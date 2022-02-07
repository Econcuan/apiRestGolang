package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
)

var movies = Movies{
	Movie{"Sin Limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo Gas", 2005, "Juan Antonio"}, 
}

func Index(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola Mundo desde mi servidor web con GO")
	}

	func Movielist(w http.ResponseWriter, r *http.Request){

		//fmt.Fprintf(w, "Listado de Peliculas")

		json.NewEncoder(w).Encode(movies)
		
	}

	func MovieShow(w http.ResponseWriter, r *http.Request){
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







