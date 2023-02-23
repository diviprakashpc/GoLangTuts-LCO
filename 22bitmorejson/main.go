package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	//Aliases
	Name     string `json:"coursename"` //Now it will create the json with key coursename from this struct.
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // dont show it
	Tags     []string `json:"tags,omitempty"` //omit nil values
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJSON()
	DecodeJson()
}

func EncodeJSON() {
	lcoCourses := []course{
		{"ReactJSBootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERNBootcamp", 699, "LearnCodeOnline.in", "abc123", []string{"full-stack-dev", "js"}},
		{"AngularBootcamp", 299, "LearnCodeOnline.in", "abc123", nil},
	}

	//package this data as JSON data
	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "lco", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	//as you know data is received in form of bytes
	//and it needs to be changed to json or stringified
	jsonDataFromWeb := []byte(`{     
	   "coursename":"ReactJS Bootcamp",
	   "Price":299,
	   "website":"LearnCodeOnline.in",
	   "tags":["web-dev","js"]
   }`) //Suppose. Calling an api and then receiveing it will be hectic. So create a sample and assume it is recieved.
	var lcoCourse course

	// checking if djson data is valid or not.

	checkValid := json.Valid(jsonDataFromWeb) //check if recieved json data is valid

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		// simply writing lcoCourse may give a copy of it and it will be problemtic
		//if we want to store some data in orignal.
		fmt.Printf("%#v\n", lcoCourse) //to print the struct/interfaces we have pound sign
	} else {
		fmt.Println("JSON was not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	//interface because we dont know the comming value type.

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", key, value, value)
	}

}
