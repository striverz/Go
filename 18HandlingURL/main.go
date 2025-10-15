package main

import (
	"fmt"
	"net/url"
)


const myUrl string="https://og.dev/learn?status=Active&search=Manikanta"
func main(){

	 fmt.Println("URL Handling in Go")

	 //exploring url in order to explore you have to parse first
	 result,_:=url.Parse(myUrl)

	 fmt.Println(result)
	 fmt.Println(result.Scheme) //https
	 fmt.Println(result.Host) //go.dev
	 fmt.Println(result.Path) // learn
	 fmt.Println(result.RawQuery) //status=Active&search=Manikanta

	 qparameters:=result.Query()
	 
	 for key,val:=range qparameters{
		fmt.Println(key,val)

	 }


	 //Creation the own URl

	 creationUrl:=&url.URL{
		Scheme: "http",
		Host:"mani.in",
		Path: "/code",
		RawPath: "user=karthik",
	 }

	 finalUrl:=creationUrl.String()

	 fmt.Println(finalUrl) //http://mani.in/code
}