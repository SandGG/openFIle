package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var file, err = os.Open("file.txt")
	var content, errC = ioutil.ReadAll(file)
	defer file.Close()

	if err != nil {
		panic("File not found")
	}
	if errC != nil {
		panic(err)
	}

	fmt.Println("File.txt content:", string(content))
	fmt.Println("File closed successful")
}
