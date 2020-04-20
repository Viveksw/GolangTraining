package main

import (
	"fmt"
	"net/http"
)

func main() {

	// create slice
	studentRoll := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(len(studentRoll))
	studentRoll[5] = 10

	//multiple defer statements
	defer fmt.Println("End of function")
	defer func(slice []int, s int) []int {
		fmt.Println("Elements removed from slices")
		return append(slice[:s], slice[s+1:]...)

	}(studentRoll, 0)

	var newSlice = make([]int, len(studentRoll))
	copy(newSlice, studentRoll)
	newSlice[6] = 15
	fmt.Println(newSlice)

	// map: create
	student := map[int]string{
		10: "Vivek",
		14: "Mayur",
		15: "XYZ",
	}
	//map: print
	fmt.Println(student)
	//map: update
	student[15] = "Human"
	//map: delete
	delete(student, 10)
	fmt.Println(student)

	//Http get request
	resp, err := http.Get("http://google.com/")

	//http get response
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Date: " + resp.Header.Get("Date"))
	}

}

/*

1. create slice with any data type
print length
update 5th element
create different copy of slice

2. map:
	create map for studewnt record, key=roll_num, value = name
	print each key, value
	modify name by giving roll number
	remove student from map

3. use multiple differ statements

4. http get request, take response
*/
