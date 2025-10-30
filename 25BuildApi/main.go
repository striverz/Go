package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Course Model
type Course struct{
	CourseId string `json:"course_id"`
	CourseName string `json:"course_name"`
	CoursePrice int `json:"course_price"`
	Author *Author `json:"author"`
}

type Author struct{
	FullName string `json:"full_name"`
	Website string `json:"website"`
}


//fake DB
var courses []Course

//Middleware or helperFunction
func (c *Course) IsEmpty() bool{
	return c.CourseId=="" && c.CourseName==""
}

func main(){
	fmt.Println("Building API In Go")

	r:=mux.NewRouter()

	r.HandleFunc("/",serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":6969",r))
}

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Created My First Go Lang Server... At PORT: 6969</h1>"))

}


func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get All Courses API")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}