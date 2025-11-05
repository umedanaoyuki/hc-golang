package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)
  
type User struct {
	Age int `json:"age"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type LogData struct {
	User User `json:"user"`
	Dist string `json:"dist"`
	Level string `json:"level"`
	Msg string `json:"msg"`
	Src string `json:"src"`
	Time string `json:"time"`
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

	file , err := os.Open(logFile.Name())
	if err != nil {
		log.Fatal("ファイルを開けませんでした", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
		
	for scanner.Scan() {
		line := scanner.Text()
		// テキスト一行ごとの処理		
		// fmt.Println(line)

		var logDatas LogData
		if err := json.Unmarshal([]byte(line), &logDatas); 
		err != nil {
			log.Printf("JSONパースエラー: %v", err)
			continue
		}
		// %+v を使用するとフィールド名も表示される
		fmt.Printf("%+v\n", logDatas)
	}

	// 関数が終了した際に確実に閉じるようにする
	defer file.Close()

	connStr := "user=test-user password=test-pass dbname=users sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	defer db.Close()
	fmt.Println("データベース接続成功")
}