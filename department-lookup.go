package main

/*
In this exercise you will create an array of employee structs, each containing a lookup code that maps into a table of department names.
You will generate a report of the employee name and department.
*/

import "fmt"

func main() {
	/*
		Create a type of struct named employee
		add two string fields
		- departmentCode
		- lname
	*/
	type employee struct {
		departmentCode string
		lname          string
	}

	/*
		Create an array of several (4 or more) employee instances
		For the department code, use the letters 't', 'e', 'a'
	*/
	var empArray [5]employee

	gussy := employee{
		departmentCode: "t",
		lname:          "gussy",
	}
	tinku := &employee{"e", "tinku"}
	pilli := &employee{"a", "pilli"}

	empArray[0] = gussy
	empArray[1] = *tinku
	empArray[2] = *pilli
	empArray[3].lname = "billi"
	empArray[4].lname = "munni"

	/*
		Assign a variable named departmentMap to a map of string key/ string value
		set the key "a" to the value "administrative"
		set the key "e" to the value "executive"
		do not assign an entry to the 't' departmentCode
	*/
	departmentMap := map[string]string{
		"a": "administrative",
		"e": "executive",
	}

	/*
		for each employee:
		if departmentCode is in map print employee name, department name
		else print employee name, "department not found"
	*/
	for _, emp := range empArray {
		description, ok := departmentMap[emp.departmentCode]
		if !ok {
			fmt.Printf("%s department not found\n", emp.lname)
		} else {
			fmt.Println(emp.lname, description)
		}
	}

}
