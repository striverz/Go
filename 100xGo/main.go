package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/striverz/100xGo/router"
)

func main(){
	fmt.Println("100x Go...")

	r:=router.Router()
	log.Fatal(http.ListenAndServe(":6969",r))

	fmt.Println("Server running at port 6969...")


}