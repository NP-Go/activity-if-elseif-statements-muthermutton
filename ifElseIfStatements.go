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
		if number < 9 && number > 0 {
			fmt.Println("The number has only 1 digit")
		} else {
			fmt.Println("The number has more than 1 digit")
		}
	} else if number%2 == 1 {
		resultMessage := strconv.Itoa(number) + " is odd"
		fmt.Println(resultMessage)
		if number < 9 && number > 0 {
			fmt.Println("The number has only 1 digit")
		} else {
			fmt.Println("The number has more than 1 digit")
		}
	} else {
		fmt.Println("This is not a number")
	}
}
