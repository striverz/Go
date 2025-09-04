package main

import (
	"fmt"
	"sort"
)

func main(){

	var numArray=[]int{1,2,3,4} //declaring array
	fmt.Printf("type of numArray %T\n",numArray)
	fmt.Println(numArray)

	//adding elements in the array
	numArray=append(numArray,100,200,300)
	fmt.Println(numArray)

	//slicing elements in the array
	numArray=append(numArray[1:])
	fmt.Println(numArray)

	numArray=append(numArray[:4])
	fmt.Println(numArray)


	//Another way of declaring the array
	scoresArray:=make([]int,4)
	scoresArray=append(scoresArray, 9,7,3,2,1,7,4,5,9); //after length 4 this append re-sizes the array to 4+(9)=13 length
	fmt.Println(scoresArray) //[0 0 0 0 1 2 3 4 5 6 7 8 9]
	
	//some useFul methods
	sort.Ints(scoresArray) //sorts the array
	fmt.Println(scoresArray)
	fmt.Println(sort.IntsAreSorted(scoresArray)) //tells wheter array sorted or not 



	//removing the specific index values from the slices(array)
	var courses=[]string{"js","ts","python","c++","java"}
	var index int=2
	fmt.Println(courses)

	courses=append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}