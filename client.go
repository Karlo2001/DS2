package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func clientGetCourses() {
	resp, err := http.Get("http://localhost:8080/courses")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))              // convert to string before print
}

func clinetGetCourseById(id string) {
	resp, err := http.Get("http://localhost:8080/courses/" + id)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))              // convert to string before print
}
