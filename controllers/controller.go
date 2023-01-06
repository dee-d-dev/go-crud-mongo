package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/dee-d-dev/go-mongodb-crud/models"
	"github.com/dee-d-dev/go-mongodb-crud/helpers"
	"github.com/gorilla/mux"
)


func CreateOne(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	result := helpers.CreateTodo(todo)

	json.NewEncoder(w).Encode(result)
}


func UpdateTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	result := helpers.UpdateTodo(params["id"])
	json.NewEncoder(w).Encode(result)

}

func Delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result := helpers.DeleteTodo(params["id"])
	json.NewEncoder(w).Encode(result)
}

func Todos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	result := helpers.GetAllTodos()
	json.NewEncoder(w).Encode(result)
}

func DeleteEverything(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

    result := helpers.DeleteAll()
    json.NewEncoder(w).Encode(result)
}