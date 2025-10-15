package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url="https://filmygpt.site/"

func main(){

	fmt.Println("Web Requests in Go")

	response,err:=http.Get(url) //makes web request to get the data

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()  //Make sure to close the Request for data leak

	data,err:=ioutil.ReadAll(response.Body) //Reads all web data in the form of 0's and 1's
	
	if err!=nil{
		panic(err)
	}

	content:=string(data) //converts the 0's and 1's data into the readable strings
	fmt.Println(content)

	fmt.Println("Web Requests Finished")



}