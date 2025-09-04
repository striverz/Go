package main

import "fmt"

func main(){

	number:=5

    //Method-01
	if number%3==0 && number%5==0{
		fmt.Println("FizzBuzz")
	}else if(number%3==0){
		fmt.Println("Fizz")
	}else if(number%5==0){
		fmt.Println("Buzz")
	}else{
		fmt.Println("Invalid")
	}


	//Method-02
	if num:=3; num<10{   //declaring & initializing here 
		fmt.Println("Number is less than 10")
	}else {
		fmt.Println("Number is greater than 10")
	}

	
}