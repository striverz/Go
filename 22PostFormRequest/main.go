package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main(){
	fmt.Println("Post Form Request in Go")

    PerformPostFormRequest()
}

func PerformPostFormRequest(){
	const myUrl="http://localhost:6969/postform"


	data:=url.Values{}
	data.Add("firstName","Manikanta")
	data.Add("lastName","Korimilli")
	data.Add("email","mani@gmail.com")

	response,err:=http.PostForm(myUrl,data)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()

	content,_:=io.ReadAll(response.Body)

	fmt.Println(string(content))

}