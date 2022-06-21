package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Estructura de los End-Points de la Appi Rest
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

//Inicializando la estructura Route
type Routes []Route

//Configuracion del MidleWare
func NewRouter() *mux.Router {
	//Llamado de libreria mux para response de la Api

	router := mux.NewRouter().StrictSlash(true)
	//Buscar las rutas disponibels para el response y ejecución de la fución
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router

}

//Urls disponebles
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Movielist",
		"GET",
		"/peliculas",
		Movielist,
	},
	Route{
		"MovieShow",
		"GET",
		"/pelicula/{id}",
		MovieShow,
	},
	Route{
		"MovieAdd",
		"POST",
		"/pelicula",
		MovieAdd,
	},
	Route{
		"MovieUpdate",
		"PUT",
		"/pelicula/{id}",
		MovieUpdate,
	},
	Route{
		"MovieRemove",
		"DELETE",
		"/pelicula/{id}",
		MovieRemove,
	},
}
	