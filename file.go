package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	data := []byte("Message text...")
	err := ioutil.WriteFile("text.txt", data, 0600)
	if err != nil {
		fmt.Println("Cannot create file")
	} else {
		fmt.Println("Ok")
	}

	file_data, ok := ioutil.ReadFile("text.txt")
	if ok != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(string(file_data))
	}
}
