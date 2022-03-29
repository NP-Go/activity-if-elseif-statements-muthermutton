package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert your code here
	//Hint if a:= ?? ; condition {  }
	if number := 17; number%2 == 0 {
		resultMessage := strconv.Itoa(number) + " is even"
		fmt.Println(resultMessage)
	} else if number%2 == 1 {
		resultMessage := strconv.Itoa(number) + " is odd"
		fmt.Println(resultMessage)
	}
}
