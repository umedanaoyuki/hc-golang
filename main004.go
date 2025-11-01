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

	for i, _ := range users {
		// rangeの中でusersの要素を変更して、もとのusersも変更させる
		users[i].Age = 44
	}

	fmt.Printf("%v", users)
}