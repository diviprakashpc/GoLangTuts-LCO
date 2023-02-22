package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This is a Files tutorial`")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, error := os.Create("./mylcogofile.txt")

	// if error != nil {
	// 	panic(error)
	// // }
	checkNilErr(error)
	length, error := io.WriteString(file, content)
	// if error != nil {
	// 	panic(error)
	// }
	checkNilErr(error)
	fmt.Println("Length is", length)
	defer file.Close()
	ReadFile("./mylcogofile.txt")
}

func ReadFile(filname string) {
	databyte, error := ioutil.ReadFile(filname) // Reading a files occure not in string format but in byte format
	checkNilErr(error)
	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
