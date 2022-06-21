package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}
	func responseMovie(w http.ResponseWriter, status int, results Movie ){
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)

	}

	func responseMovies(w http.ResponseWriter, status int, results []Movie ){
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)

	}

var collection = getSession().DB("curso_go").C("movies")

//Función master para listado o view de movies
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde mi servidor web con GO")
}

//Listando las peliculas
func Movielist(w http.ResponseWriter, r *http.Request) {

	var results []Movie
	fmt.Println(results)
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados:", results)
	}

	responseMovies(w,200, results)
}

//Mostrando las peliculas

func MovieShow(w http.ResponseWriter, r *http.Request) {
	
	//Conectando a base de datos para ver las peliculas

	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if err != nil{
		w.WriteHeader(404)
		return 
	}

	responseMovie(w,200, results)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
			return
	}

	responseMovie(w,200, movie_data)
	
}
	
func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	//Codigo para modificar pelicula existente
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id){
		w.WriteHeader(404)
		return
	}
		oid := bson.ObjectIdHex(movie_id)

		decoder := json.NewDecoder(r.Body)

		var movie_data Movie
		err := decoder.Decode(&movie_data)

		if err != nil {
			panic(err)
			w.WriteHeader(500)
			return
		}

		defer r.Body.Close()

		document  := bson.M{"_id": oid}
		change := bson.M{"$set": movie_data}
		err = collection.Update(document, change)

		if err != nil {
		panic(err)
		w.WriteHeader(400)
		return
		}

	responseMovie(w,200, movie_data)
}

type Message struct {
	Status string `json:"Status"`
	Message string `json:"Message"`
}

	func (this *Message) setStatus(data string ){
		this.Status = data
	}

	func (this *Message) setMessage(data string ){
		this.Message = data
	}

func MovieRemove(w http.ResponseWriter, r *http.Request) {
	//para eliminar peliculas de la lista
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	err := collection.RemoveId(oid)

	if err != nil{
		w.WriteHeader(404)
		return
	}

		//results := Message{"succses", "La pelicula con ID"+movie_id+"ha sido borrada correctamente"}

		message := new(Message)

		message.setStatus("success")
		message.setMessage("La pelicula con ID"+movie_id+"ha sido borrada correctamente")

		results := message
		w.Header().Set("content-Type","application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(results)
}




