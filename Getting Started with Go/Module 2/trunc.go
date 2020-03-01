package main

/* This program gets a float number from user input
   and convert it to an integer. Output the result */

import "fmt"

func trunc() {
	var num float32
	fmt.Printf("Please input a float number:")
	fmt.Scan(&num)
	fmt.Println("The truncated numner is", int(num))
}

func main() {
	trunc()
}

/*
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Input a number: ")
	var input string
	fmt.Scan(&input)
	parsed, _ := strconv.ParseFloat(input, 64)
	fmt.Printf("truncated number is %d\n", int(parsed))
}*/
