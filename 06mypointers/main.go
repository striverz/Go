package main

import "fmt"

func main(){
	msg:="Welcome to Pointers"
	fmt.Println(msg);
  
	var ptr *int;
	fmt.Println("value is ",ptr); //value is nil

	myNumber:=10

	var ptr2 *int=&myNumber
	fmt.Println("value is ",ptr2) //address xxfdfdfd
	fmt.Println("value is ",*ptr2) //value 10
	fmt.Println("value is ",&myNumber) //address xxfdfd
	fmt.Println("value is ",&ptr2) //new address zzdfdfd

	//useCase or Expo

	*ptr2=*ptr2+10
	fmt.Println("value is ",myNumber)


	
}