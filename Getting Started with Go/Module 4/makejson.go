/*
	Write a program which prompts the user to first enter a name, 
	and then enter an address. Your program should create a map 
	and add the name and address to the map using the keys “name” 
	and “address”, respectively. Your program should use Marshal() 
	to create a JSON object from the map, and then your program 
	should print the JSON object.
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

func makejson() {
	var name, address string
	fmt.Println("Please enter your name:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Println("Please enter your address:")
	scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		address = scanner.Text()
	}
	m := map[string]string{"name": name, "address": address}
	barr, err := json.Marshal(m)
	if err == nil {
		fmt.Printf("%s\n", string(barr))
	}
}

func main() {
	makejson()
}
