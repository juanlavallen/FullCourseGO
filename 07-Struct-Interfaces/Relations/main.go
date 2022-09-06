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

func main() {

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
}
