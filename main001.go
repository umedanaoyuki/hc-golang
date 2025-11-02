package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	// 文法書的にはinterface{} を使用すると思うのですが、最新のGoのバージョンではany型が使用できるらしいので使用しました
	n := []any{1, "2", 10, "11"}
	// 文字列を格納するためのスライス
	stringN := []string{}

	for _, v := range n {
		switch v := v.(type) {
		// int型の場合は文字列に変換
		case int:
			stringN = append(stringN, strconv.Itoa(v))
		case string:
			stringN = append(stringN, v)
		}
	}

	for _, z := range stringN {
		// 文字数カウント
		count := utf8.RuneCountInString(z)
		if count == 2 {
			fmt.Println(z)
		} else {
			fmt.Println("0" + z)
		}	
	}
}


