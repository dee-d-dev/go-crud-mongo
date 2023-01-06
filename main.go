package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/dee-d-dev/go-mongodb-crud/routers"
)

func main() {

	fmt.Println("Welcome")
	fmt.Println("Connecting to server")

	r := routers.Router()

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("server running on 4000")

}
