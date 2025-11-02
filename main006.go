package main

import "fmt"

type MyInt int

// MyInt型に関連したStringのメソッドを追加して、そこで""hoge"を返却するようにする
func (m MyInt) String() string {
	return "hoge"
}

func main() {
	var m MyInt = 3

	fmt.Println(m)
}