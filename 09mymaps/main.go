package main

import "fmt"

func main(){

	userDetails:=make(map[int]string)  //defining map

	userDetails[1]="Manikanta"  //adding values in map
	userDetails[2]="Karthik"
	userDetails[3]="Panee"
	userDetails[4]="Saranya"

	fmt.Println("The Map is: ",userDetails)

	fmt.Println(userDetails[2]) //accessing element

	delete(userDetails,3) //deleting element

	fmt.Println(userDetails)



	//looping through map

	for key,value :=range userDetails{

		fmt.Printf("The key & Pairs values are %v --> %v\n",key,value)
	}




}