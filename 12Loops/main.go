package main

import "fmt"

func main(){

	days:=[]string{"Sunday","Monday","Tuesday","Wednesday","Thrusday"}
	fmt.Println(days)

	//Loop-01
	for i:=0; i<len(days); i++{
		fmt.Printf("%v : %v\n",i,days[i])
	}

	//Loop-02
	for j:=range days{
		fmt.Println(j)  //prints the index
	}

	//Loop-03 (for Each loop)
	for index,value:=range days{
		fmt.Println(index,value)
	}


	//Now, this loop behaves like while loop ok
	value:=1
	for value<10{
		if value==2{
			break
		}
		fmt.Println(value)
		value++
	}

	fmt.Println("GoTo")

	newValue:=1
	for newValue<10{
		if(newValue==3){
			goto leo
		}
		fmt.Println(newValue)
		newValue++
	}

	leo:
	   fmt.Println("This is Leo..")

	
	


}