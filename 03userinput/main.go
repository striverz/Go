package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome:="Welcome to the Restaurant"
	fmt.Println(welcome)

	reader:=bufio.NewReader(os.Stdin);
	fmt.Println("Enter rating for the pizza");

	//comma ok || error ok syntax
	input,_:=reader.ReadString('\n');
	fmt.Println("Thank you for rating: ",input)
	fmt.Printf("The type of reader is: %T\n",input);

	
}
