package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Working with Files...")

	content:="This is Manikanta Korimilli"

	//creating File 
	file,err:=os.Create("./myFile.txt")

	if err!=nil{
		panic(err)
	}

	//Putting data into that file 
	length,err:=io.WriteString(file,content)

	if err!=nil{
		panic(err)
	}
	fmt.Println("length of the file is: ",length)
    
	//reading the File
	readFile("./myFile.txt")
}

func readFile(filePath string){

	data,err:=ioutil.ReadFile(filePath)
	if err!=nil{
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))
}