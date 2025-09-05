package main

import "fmt"

func main(){

	userOne:=User{"Manikanta","manikanta@gmail.com",true,21}
	fmt.Println(userOne.Name)
	fmt.Println(userOne.Email)


	//Using Methods
	userOne.GetStatus()

	userOne.GetNewMail()

}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

//Defining Method

func (u User) GetStatus(){     //This is Method Signature of defining...
	fmt.Println("The status is: ",u.Status)
}


func (u User) GetNewMail(){
	u.Email="sara@gmail.com"
	fmt.Println(u.Email)

}