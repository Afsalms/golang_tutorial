package main

import (
	"os"
	"fmt"
)



func main(){
	path := "test.txt"
	file, err := os.Open(path)
	if err != nil{
		fmt.Println("Error on opening  file %s", err)
	}
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("Error while reading file info %s", err)
	}
	fmt.Println(string(fi))
	file.Close()

}
