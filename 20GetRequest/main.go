package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET Request in Go Lang")
	PerformGetRequest()
}

func PerformGetRequest() {
	
	const url = "http://localhost:6969/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response code: ",response.StatusCode)

	content, _ := io.ReadAll(response.Body)

	//case:01
	fmt.Println(string(content))

	//case:02
	var responseString strings.Builder
	byteCount,_:=responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	

}
