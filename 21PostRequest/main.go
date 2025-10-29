package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func main(){
	fmt.Println("Post Request in Go Lang")

	PerformPostRequest();
}


func PerformPostRequest(){
	const url="http://localhost:6969/post"

    
	//Making Dummy Json Data
	responseBody:=strings.NewReader(`
			{
	         "name":"Manikanta",
			 "age":"21",
			 "married":"false"
			}
	`)	

	response,err:=http.Post(url,"application/json",responseBody)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close();

    //diff things
	fmt.Println(response.StatusCode)
	fmt.Println(response.Request.Method)
	fmt.Println(response.Request.URL)
	fmt.Println(response.ContentLength)
	fmt.Println(response.Status)
	
	

	content,_:=io.ReadAll(response.Body)

	//Request Body 
	fmt.Println(string(content))
	 
}