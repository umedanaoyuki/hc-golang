package main

import (
	"fmt"
	"slices"
)


type MyIntSlice []int

func (m MyIntSlice) Unique() MyIntSlice {
	// 重複を削除したスライス
	n := []int{}
	
	for _, v := range m {
		// スライスの中にvの値がすでに含まれているかどうかのチェック
		result := slices.Contains(n, v)
		if !result {
			n = append(n, v)
		}
	}
	return n
}

func main() {
	m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Println(m.Unique())
}