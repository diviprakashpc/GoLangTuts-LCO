package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Verb video")

	//PerformGetRequest()
	//PerformPostJsonRequest()
	//PerformPostFormDataRequest()
}

func PerformPostJsonRequest() {
	//Usually we send post data in form of json or form-url-encoded
	const myurl = "http://localhost:1111/post"
	//fake json payload
	requestBody := strings.NewReader(`
    {
		"courseName":"Let's go with golang",
		"price":0,
		"platform":"LearnCodeOnline.in"
	}
`)

	response, err := http.Post(myurl, "application/json", requestBody) // url , content-type , request-body

	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(content))

}

func PerformPostFormDataRequest() {
	const myurl = "http://localhost:1111/postform"
	//formdata

	data := url.Values{}
	fmt.Println(data) // it is a map
	data.Add("firstname", "divya")
	data.Add("lastname", "Prakash")
	data.Add("email", "diviprakash3@gmail.com")

	response, err := http.PostForm(myurl, data) //Dont need to tell content type now'

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	fmt.Println(string(content))

}

func PerformGetRequest() { //Public for upper case
	const myurl = "http://localhost:1111/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Current length is : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	dataInString := responseString.String()
	fmt.Println("ByteCount is ", byteCount)
	fmt.Println(dataInString)
	//This content is in data byte format

	//fmt.Println("The recieved response is : ", string(content))

}
