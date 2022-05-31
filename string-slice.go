package main

/*
In this exercise you will take the following text and parse it in various ways:
Today it is the beginning of spring. Let's rejoice, the sun is out!
Assign the above text to a variable.
Using just the Index and Trim functions found in the strings package, implement a program which outputs the following (you can assume the period terminates the first sentence and that spaces separate words):
The first word ends at index 5 and is 'Today'
The last sentence is: 'Let's rejoice, the sun is out!'
The last six characters of the string are: 's out!'
The first five characters of the second sentence are: 'Let's'
*/

import (
	"fmt"
	"strings"
)

func main() {
	s := "Today it is the beginning of spring. Let's rejoice, the sun is out!"
	firstSpace := strings.Index(s, " ")
	lastDot := strings.LastIndex(s, ".")
	fmt.Printf("The first word ends at index %d and is '%s'\n", firstSpace, s[:firstSpace])
	fmt.Printf("The last sentence is: '%s'\n", strings.Trim(s[lastDot+1:], " "))
	fmt.Printf("The last six characters of the string are: '%s'\n", s[len(s)-6:])
	fmt.Printf("The first five characters of the second sentence are: '%s'\n", strings.Trim(s[lastDot+1:], " ")[:5])
}
