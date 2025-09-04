package main

import "fmt"

func main(){
	msg:="Welcome to Arrays Cls"
	fmt.Println(msg)
	
	var fruitList [4]string;

	fruitList[0]="Apple"
	fruitList[1]="Banana"
	fruitList[3]="Grapes"

	fmt.Println("Array value is: ",fruitList)  //[Apple Banana    Grapes] --> little bit space in index2
	fmt.Println("Array length is: ",len(fruitList)) //length is 4


	var vegList=[3]string{"Carrot","Brinjal","Potato"}

	fmt.Println("Array value is: ",vegList)
	fmt.Println("Array length is: ",len(vegList))
	fmt.Println("Arr value index is: ",vegList[1]);






}