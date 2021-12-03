package main

import (
	"fmt"
  "io/ioutil"
  "strconv"
)

func main() {
	urlFile := "../../inputzz/day1.txt"
  var result []int

	byteRead, err := ioutil.ReadFile(urlFile)

	if err != nil {
		fmt.Printf("Error leyendo el archivo: %v", err)
	}

  for _, byteR := range byteRead {
    data := string(byteR)
    i, _ := strconv.Atoi(data)
    fmt.Printf("%v", i)
  }


	//var count int = 0

	/*for _, number := range content {
		if number > firstNumber {
			count++
			firstNumber = number
		} else {
      firstNumber = content[0]
    }
	}*/
}
