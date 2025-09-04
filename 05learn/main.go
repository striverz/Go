package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("Learning about Time in GO")

	//printingt the current time in diff formats
	presentTime:=time.Now();
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))


	//creating the time
	createdDate:=time.Date(2025,time.January,4,4,2,1,3,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))



	
}