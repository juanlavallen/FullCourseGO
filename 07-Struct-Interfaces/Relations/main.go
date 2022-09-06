package main

import "fmt"

// One to One

type User struct {
	name   string
	email  string
	status bool
}

type Student struct {
	user User
	code string
}

// One to Many

type Course struct {
	title  string
	videos []Video
}

type Video struct {
	title    string
	duration int
	course   Course
}

func main() {

	// One to One

	user := User{
		name:   "Bill",
		email:  "bill@bill.net",
		status: true,
	}

	student := Student{
		user: user,
		code: "0441975",
	}

	fmt.Println(student)            // {{Bill bill@bill.net true} 0441975}
	fmt.Println(student.user)       // {Bill bill@bill.net true}
	fmt.Println(student.user.email) // bill@bill.net

	// One to Many

	video1 := Video{title: "Introduction", duration: 5000}
	video2 := Video{title: "Data Structure", duration: 5000}

	course := Course{
		title:  "Full Course GO",
		videos: []Video{video1, video2},
	}

	fmt.Println(course)
}
