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


	responseBody:=strings.NewReader(`
			{
	         "name":"Manikanta",
			 "age":"21"

			
			}
	`)	

	response,err:=http.Post(url,"application/json",responseBody)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close();

	content,_:=io.ReadAll(response.Body)

	fmt.Println(string(content))
}