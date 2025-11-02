package main

import (
	"errors"
	"fmt"
)

func findKeyByValue(m map[int]string, v string) (int, error) {
	for i, z := range m {
		if z == v {
			return i, nil
		}
	}
	return 0, errors.New("エラー")
}

func main() {
	m := map[int]string{
	  1: "01",
	  2: "02",
	  3: "03",
	}

    // ここの記述のコメントアウトを有効＆解除することででバック可能
	key, err := findKeyByValue(m, "03")
	// key, err := findKeyByValue(m, "05")
	if err != nil {
		fmt.Println(key, err)
	} else {
		fmt.Println(key, err)
	}
  }