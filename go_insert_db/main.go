package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//  {
// 	"user": {
// 	  "age": 22,
// 	  "name": "tarou",
// 	  "role": "student"
// 	},
// 	"dist": "PostgreSQL",
// 	"level": "info",
// 	"msg": "test",
// 	"src": "backend",
// 	"time": "2021-08-01T00:05:05Z"
  
type User struct {
	age int
	name string
	role string
}

type LogData struct {
	user User
	dist string
	level string
	msg string
	src string
	time string
 }

 type Person struct {
	Name   string
	Age    int
	gender string
}

func main() {
	//実行時に引数を受け取る
	// ファイル名の指定
	args := os.Args
	// ファイルは一つだけと限定する
	if len(args) != 2 {
		fmt.Println("引数の数が間違っています")
		// プロセス終了（失敗）
		os.Exit(1)
	}

	fmt.Println("ファイル読み取り処理を開始します")
	// ファイルをOpenする
	logFile, err := os.Open(args[1])
	// 読み取り時の例外処理
	if err != nil {
		fmt.Println("error")
	}

	// p := Person{
	// 	Name: "Mike",
	// 	Age:  20,
	// 	gender: "male",
	// }

	file , err := os.Open(logFile.Name())
	// scanner := bufio.NewScanner(file)
	if err != nil {
		log.Fatal("ファイルを開けませんでした", err)
	}

	// fmt.Println(file)

	defer file.Close()

	scanner := bufio.NewScanner(file)
		
	for scanner.Scan() {
		line := scanner.Text()
		// テキスト一行ごとの処理		
		fmt.Println(line)
	}


	// }
	// err = scanner.Err()   
	// if err != nil {
	// 	fmt.Println("error")
	// 	// プロセス終了（失敗）
	// 	os.Exit(1)
	// }

	// m, _ := json.Marshal(p)
	// fmt.Println(string(m))

	// 関数が終了した際に確実に閉じるようにする
	defer file.Close()
}