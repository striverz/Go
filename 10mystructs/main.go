package main

import "fmt"

//Defining struct
type User struct{
	Name string
	Email string
	Status bool
	Age int
}
func main(){

	manikanta:=User{"Manikanta","manikanta@gmail.com",true,21}

	fmt.Printf("The User details are: %+v\n",manikanta)
	fmt.Println("Name: ",manikanta.Name)
	fmt.Println("Email: ",manikanta.Email)
	fmt.Println("Status: ",manikanta.Status)
	fmt.Println("Age: ",manikanta.Age);
	
}


