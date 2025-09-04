package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
   message:="Welcome to the learning of Go"
   fmt.Println(message)

   reader:=bufio.NewReader(os.Stdin);
   input,_:=reader.ReadString('\n')

  input=strings.TrimSpace(input);
  convertedOne,err:=strconv.ParseFloat(input,64);

  if err!=nil{
	fmt.Println(err)
  }else{
	fmt.Println(convertedOne+1)
	fmt.Printf("%T\n",convertedOne)
  }

   
}