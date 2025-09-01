package main

import "fmt"

const apiKey="69:69"
func main(){

	//method-01
	var name string="Manikanta Korimilli";
	fmt.Println(name)
	fmt.Printf("The type of variable: %T \n",name);

	//method-02
	var totalAmount=40
	fmt.Println(totalAmount);
	fmt.Printf("Type type of variable: %T \n",totalAmount)

	//method-03
	isLoggin:=true //valorous operator (NOTE:works only inside the method)
	fmt.Println(isLoggin)
	fmt.Printf("The type of variable: %T \n",isLoggin)

	fmt.Println(apiKey)
	fmt.Printf("The type of variable: %T \n",apiKey)

	var defaultValue int;
	fmt.Println(defaultValue);

}