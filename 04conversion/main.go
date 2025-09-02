package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to the pizza app")
	fmt.Println("Enter Rating")
	
	reader:=bufio.NewReader(os.Stdin)
	input,_:=reader.ReadString('\n')
	fmt.Println("Your Rating is: ",input)

	input=strings.TrimSpace(input)

	numberInput,err:=strconv.ParseFloat(input,64);

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Increased Rating is: ",numberInput+1);
		fmt.Printf("The type is %T\n",numberInput)
	}

}