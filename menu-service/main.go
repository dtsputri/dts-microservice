package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dtsputri/dts-microservice/menu-service/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))
	fmt.Println("menu service listen at port 3000")
	log.Panic(http.ListenAndServe(":3000", router))
}
