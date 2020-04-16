package main

import (
	"fmt"
	"os"
	"strconv"
)

// ContactInfo struct
type ContactInfo struct {
	phoneNumber string
	address     string
}

// Student struct
type Student struct {
	id      int
	name    string
	age     int
	contact ContactInfo
}

// update Student
func (s *Student) updateStudent(newStudent Student) {

	s.name = newStudent.name
	s.id = newStudent.id
	s.age = newStudent.age
	s.contact.phoneNumber = newStudent.contact.phoneNumber
	s.contact.address = newStudent.contact.address
}

func (s *Student) getAge() int {
	return s.age
}

func (s Student) toString() string {
	id := strconv.Itoa(s.id)
	age := strconv.Itoa(s.age)

	return (id + ", " + s.name + ", " + age + ", " + s.contact.phoneNumber + ", " + s.contact.address)
}

// Write student to File
func writeToFile(stud Student) {
	f, err := os.Create("./studentinfo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(stud.toString())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "info written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	student := Student{19, "Mayur", 27, ContactInfo{"1234", "Calsoft Pune"}}
	student.updateStudent(Student{19, "Vivek", 28, ContactInfo{"456", "Calsoft Pune"}})
	fmt.Println(student.getAge())
	writeToFile(student)
}

/*
1.
struct for Student
nesting of Struct ContactInfo

print student info


2.Update info of student


3. test cases for checking contact details of student and age

4. save student info into file
*/
