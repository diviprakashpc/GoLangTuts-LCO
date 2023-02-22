package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456"

func main() {
	fmt.Println("Welcome to Tutorial on urls in Go")
	fmt.Println(myurl)

	//parsing the url
	result, error := url.Parse(myurl)
	checkNilError(error)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) //params
	//query parameters
	qparams := result.Query() //Returns the params in form of map
	fmt.Printf("The data type of qparams is %T\n", qparams)
	fmt.Println(qparams)

	//Creating a url
	partsOfUrl := &url.URL{ //We always pass a reference of url
		Scheme:  "https", //wriite these properties as it is
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=divya",
	}
	fmt.Printf("The type of partsUrl %T\n", partsOfUrl)
	anotherurl := partsOfUrl.String()
	fmt.Println(anotherurl)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
