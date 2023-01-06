package routers

import (
	"github.com/gorilla/mux"
	"github.com/dee-d-dev/go-mongodb-crud/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/todo", controllers.Todos).Methods("GET")
	router.HandleFunc("/api/todo", controllers.CreateOne).Methods("POST")
	router.HandleFunc("/api/todo/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/todo/{id}", controllers.Delete).Methods("DELETE")
	router.HandleFunc("/api/todo", controllers.DeleteEverything).Methods("DELETE")

	return router
}