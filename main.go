package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	var file, err = os.Open("./file.txt")
	var content, errC = ioutil.ReadAll(file)
	defer file.Close()

	if err != nil {
		log.fatal("File not found")
	}
	if errC != nil {
		log.fatal(err)
	}

	fmt.Println("File.txt content:", string(content))
	fmt.Println("File closed successful")
}
