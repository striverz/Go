package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const url="https://filmygpt.site"
func main(){

	fmt.Println("Learning WebRequests..")

	response,err:=http.Get(url)

	if err!=nil{
		panic(err)
	}

	fmt.Printf("Response of type %T\n",response)

	defer response.Body.Close()

	results,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	contest:=string(results)
	fmt.Println(contest)
}