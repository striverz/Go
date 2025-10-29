package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Mod in Go")
	 r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000",r))
}

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Go Lang Learning...</h1>"))

}