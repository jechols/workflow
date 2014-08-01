package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	var template string

	fileBytes, err := ioutil.ReadFile("./template.txt")
	if err != nil {
		fmt.Println("Error trying to read the template: ", err)
		os.Exit(1)
	}

	template = string(fileBytes)
	fmt.Println(template)
}
