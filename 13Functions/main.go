package main

import "fmt"

func main(){
	fmt.Println("Learning about Functions...")

	greeter();

	result:=adder(2,3)
	fmt.Println(result)

	proAdderResult:=proAdder(1,2,3,4,5,2,4,4,4,4)
	fmt.Println(proAdderResult)

	val,msg:=multiResult(1,2,3,4,5,6,6,)
	fmt.Println(val)
	fmt.Println(msg)

	


}

func greeter(){ //greeter function
	fmt.Println("Welcome")
}

func adder(valueOne int,valueTwo int) int{ //function Signature
	return valueOne+valueTwo
}

func proAdder(values ...int) int{  //...(verodic function) add unknown number of values on the Go...

	total:=0
	for _,value:=range values{
		total+=value
	}
	return total
}


func multiResult(values ...int) (int,string){ //returnting multiple results from that function signature
	total:=0
	message:="Done"

	for _,value:=range values{
		total+=value

	}
	return total,message
}