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

func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get One Course API")
	w.Header().Set("Content-Type","application/json")

	//Send Only One course which is asked by the User

	params:=mux.Vars(r)

	fmt.Println(params)

	for _,course:=range courses{

		if course.CourseId==params["id"]{

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Courses Found")
	return

}


func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create One Course API")
	w.Header().Set("Content-Type","application/json")

	//what if No Data
	if r.Body==nil{
		json.NewEncoder(w).Encode("No Data")
		return
	}

	// what if {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No Data Found Inside the Course")
		return
	}

	course.CourseId="fssdlkjf"

	courses=append(courses, course)

	json.NewEncoder(w).Encode(course)
}