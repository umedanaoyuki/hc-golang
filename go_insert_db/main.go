package main

import (
	"bufio"
	"context"
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

/**
  ログデータの構造体
*/
type LogData struct {
	User User `json:"user"`
	Dist string `json:"dist"`
	Level string `json:"level"`
	Msg string `json:"msg"`
	Src string `json:"src"`
	Time string `json:"time"`
 }

func insertData(ctx context.Context, db *sql.DB, logDatas LogData) (err error) {

	// Create a helper function for preparing failure results.
    fail := func(err error) error {
        return fmt.Errorf("error %w", err)
    }

    // Get a Tx for making transaction requests.
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return fail(err)
    }

	defer tx.Rollback()

	for _, logData := range []LogData{logDatas} {
		_, err := db.Exec("INSERT INTO users (age, name, role) VALUES ($1, $2, $3)", logData.User.Age, logData.User.Name, logData.User.Role)
		// log.Println("データ挿入成功", logData)
			if err != nil {
				log.Fatal("データ挿入失敗", err)
			}	
	}

	// fmt.Println("データ挿入成功")

	// Commit the transaction.
    if err = tx.Commit(); err != nil {
        return fail(err)
    }
	return nil
}


func main() {
	//実行時に引数を受け取る
	// ファイル名の指定
	args := os.Args
	// ファイルは一つだけと限定する
	if len(args) != 2 {
		fmt.Println("引数が間違っています。3つ以上のファイルは読み込めません。")
		// プロセス終了
		os.Exit(1)
	}

	// fmt.Println("ファイル読み取り処理を開始します")
	// ファイルをOpenする
	logFile, err := os.Open(args[1])
	// 読み取り時の例外処理
	if err != nil {
		fmt.Println("うまくファイルを読み取れませんでした", err)
		os.Exit(1)
	}

	file , err := os.Open(logFile.Name())
	if err != nil {
		log.Fatal("ファイルを開けませんでした", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// DB接続情報
	connectInfo := "user=test-user password=test-pass dbname=users sslmode=disable"

	db, err := sql.Open("postgres", connectInfo)
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	fmt.Println("データベース接続成功")
		
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
		// fmt.Printf("%+v\n", logDatas)
		// DB操作
		ctx := context.Background()
		insertData(ctx,db, logDatas)
	}

	defer db.Close()
	fmt.Println("データベース接続を終了")

}