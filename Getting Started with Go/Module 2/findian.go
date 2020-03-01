/* Write a program which prompts the user to enter a string. 
   The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. 
   The program should print “Found!” if the entered string starts with the character ‘i’, 
   ends with the character ‘n’, and contains the character ‘a’. The program should print 
   “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter 
   if the characters are upper-case or lower-case.
*/

package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func findian(str string) {
	str = strings.Trim(str, " ")
	if strings.IndexRune(str, 'i') == 0 && strings.LastIndexByte(str, 'n') == len(str)-1 && strings.ContainsRune(str, 97) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func main() {
	var str string
	fmt.Printf("Please input a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	findian(strings.ToLower(str))
}
