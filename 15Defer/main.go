package main

import "fmt"

// when we use Defer it will getExecuted or Printed at the End Culybase (}) of that function
// if there are mutliple defer statements then it works in LIFO Order
func main(){
	fmt.Println("Learning Defer...")

	fmt.Println("Hello")
	fmt.Println("World")

	defer fmt.Println("Hello") //This will Execute in the last
	fmt.Println("World")



	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello World") //Hello World Three Two One



	
}