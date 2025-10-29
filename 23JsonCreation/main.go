package main

import (
	"encoding/json"
	"fmt"
)


type course struct{
	Name string `json:"course_name"`
	Price int    `json:"price_value"`
	Platform string `json:"platform"`
	Password string `json:"-"`
	Tags []string  `json:"tags,omitempty"`
}
func main(){
	fmt.Println("Json Creation In Go")
	EncodeJson()
	DecodeJson()
}


func EncodeJson(){

	fmt.Println("Encoding of Json")

	myCourses:=[]course{
		{"tuf+",100,"Tuf","123",[]string{"coding","strive","lld"}},{"tle",50,"TLE","321",nil},
	}

	finalJson,err:=json.MarshalIndent(myCourses,"","\t")

	if err!=nil{
		panic(err)
	}

	fmt.Printf("%s\n",finalJson)



}


func DecodeJson(){


	fmt.Println("Decoding of Json")

	jsonDataFromWeb:=[]byte(`
	{
                "course_name": "tuf+",
                "price_value": 100,
                "platform": "Tuf",
                "tags": [
                        "coding",
                        "strive",
                        "lld"
                ]
        }`)
	
	var myCourses course
	checkValid:=json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("Valid Json Data")
		json.Unmarshal(jsonDataFromWeb,&myCourses)
		fmt.Printf("%#v\n",myCourses)

	}else{
		fmt.Println("Invalid Json Data")
	}

	
}