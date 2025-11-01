package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}

	for _, user := range users {
		user.Age = 44
	}

	fmt.Printf("%v", users) // どうなる？
}