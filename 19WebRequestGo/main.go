package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web request tutorial in Go")
	response, error := http.Get(url)
	checkNilErr(error)
	fmt.Printf("The type of response is : %T \n", response)
	//We are receiving the pointer to the response.i.e. *http.Response
	//So its copy is not recieved and orignal response can be accessed
	//from the recieved reference/pointer.

	fmt.Println("And the response is", response)
	/*The response object is
	    &{200 OK 200 HTTP/2.0 2 0 map[Age:[107289]
		Cache-Control: [public, max-age=0,
		must-revalidate] Content-Type:[text/html; charset=UTF-8]
		Date:[Tue, 21 Feb 2023 03:50:23 GMT] E
		tag:["b3141ad8710e6c62813315548dea5d0e-ssl-df"]
		Server:[Netlify]
		Strict-Transport-Security:[max-age=31536000] V
		ary:[Accept-Encoding]
		X-Nf-Request-Id:[01GSW7T911RDH7ZQSFPDYFY6F5]] 0xc0004be180 -1 [] false true map[] 0xc0000fe100 0xc0000e0420}

	*/
	defer response.Body.Close()                       // It is caller' responsibility to close the connection
	databytes, error := ioutil.ReadAll(response.Body) //majority of reading would be done by ioutil.
	checkNilErr(error)
	content := string(databytes)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

//WE use http module
