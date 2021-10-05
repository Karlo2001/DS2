package main

import "fmt"

var action1 int
var courseID string

func main() {
	initServer()
	for {
		fmt.Println("Type 1 to GET all courses")
		fmt.Println("Type 2 to GET course by ID")
		fmt.Scanln(&action1)
		if action1 == 1 {
			clientGetCourses()
		} else if action1 == 2 {
			fmt.Println("Type the ID of the course you would like to GET (1-2)")
			fmt.Scanln(&courseID)
			clinetGetCourseById(courseID)
		} else {
			fmt.Println("Invalid input, please try again")
		}
	}
}
