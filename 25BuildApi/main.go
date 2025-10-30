package main

import "fmt"

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
}