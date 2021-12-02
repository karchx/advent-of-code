package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	urlFile := "../inputzz/input.txt"
	byteRead, err := ioutil.ReadFile(urlFile)

	if err != nil {
		fmt.Printf("Error leyendo el archivo: %v", err)
	}

	content := string(byteRead)

	fmt.Printf("text %s", content)
}
