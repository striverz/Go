package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Course Model
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

// Fake DB
var courses = []Course{
	{
		CourseId:    "1",
		CourseName:  "Go Lang Mastery",
		CoursePrice: 299,
		Author: &Author{
			FullName: "John Doe",
			Website:  "https://johndoe.dev",
		},
	},
	{
		CourseId:    "2",
		CourseName:  "ReactJS for Beginners",
		CoursePrice: 249,
		Author: &Author{
			FullName: "Jane Smith",
			Website:  "https://janesmith.io",
		},
	},
	{
		CourseId:    "3",
		CourseName:  "Python Web Development",
		CoursePrice: 199,
		Author: &Author{
			FullName: "Michael Green",
			Website:  "https://mikegreen.tech",
		},
	},
	{
		CourseId:    "4",
		CourseName:  "Fullstack with Node.js",
		CoursePrice: 349,
		Author: &Author{
			FullName: "Alice Johnson",
			Website:  "https://alicejohnson.me",
		},
	},
	{
		CourseId:    "5",
		CourseName:  "Docker & Kubernetes Bootcamp",
		CoursePrice: 399,
		Author: &Author{
			FullName: "Robert Brown",
			Website:  "https://robertbrown.dev",
		},
	},
}

// Middleware/helper function
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Building API with Chi Router")

	r := chi.NewRouter()

	// Routes
	r.Get("/", serveHome)
	r.Get("/courses", getAllCourses)
	r.Get("/course/{id}", getOneCourse)
	r.Post("/course", createOneCourse)
	r.Put("/course/{id}", updateOneCourse)
	r.Delete("/course/{id}", deleteCourse)

	log.Fatal(http.ListenAndServe(":6969", r))
}

// Controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Created My First Go Lang Server... At PORT: 6969</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses API")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course API")
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Courses Found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course API")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No Data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data Found Inside the Course")
		return
	}

	course.CourseId = fmt.Sprintf("%d", len(courses)+1)
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course API")
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	for index, value := range courses {
		if value.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = id

			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("The Course Id is Not Found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course API")
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	for index, value := range courses {
		if value.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted")
			return
		}
	}

	json.NewEncoder(w).Encode("No Courses found with that Id")
}
