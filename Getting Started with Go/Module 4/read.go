/*
	Write a program which reads information from a file and 
	represents it in a slice of structs. Assume that there 
	is a text file which contains a series of names. Each 
	line of the text file has a first name and a last name, 
	in that order, separated by a single space on the line.

	Your program will define a name struct which has two fields,
	 fname for the first name, and lname for the last name. 
	 Each field will be a string of size 20 (characters).

	Your program should prompt the user for the name of the 
	text file. Your program will successively read each line 
	of the text file and create a struct which contains the 
	first and last names found in the file. Each struct created 
	will be added to a slice, and after all lines have been read 
	from the file, your program will have a slice containing one 
	struct for each line in the file. After reading all lines 
	from the file, your program should iterate through your slice 
	of structs and print the first and last names found in each struct.
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func read() {
	var slice []Person
	var file_name string

	fmt.Println("Please enter the name of the text file:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		file_name = scanner.Text()
	} else {
		return
	}
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File context:")
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Split(line, " ")
		var p Person
		if len(name[0]) > 20 {
			name[0] = name[0][0:20]
		}
		p.fname = name[0]
		if len(name[1]) > 20 {
			name[1] = name[1][0:20]
		}
		p.lname = name[1]
		slice = append(slice, p)
	}

	file.Close()

	for _, name := range slice {
		fmt.Println(name)
	}
}

func main() {
	read()
}