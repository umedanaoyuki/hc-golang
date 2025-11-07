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

func insertData(ctx context.Context, db *sql.DB, logData LogData, tx *sql.Tx) (err error) {

	for _, logData := range []LogData{logData} {
		_, err := db.ExecContext(ctx,"INSERT INTO users (age, name, role) VALUES ($1, $2, $3)", logData.User.Age, logData.User.Name, logData.User.Role)
			if err != nil {
				log.Fatal("データ挿入失敗", err)
			}	
	}

	fail := func(err error) error {
		return fmt.Errorf("error %w", err)
	}

    if err = tx.Commit(); err != nil {
        return fail(err)
    }
	return nil
}


func main() {
	// 実行時に引数を受け取る
	// ファイル名の指定
	args := os.Args
	// ファイルは一つだけと限定する
	if len(args) != 2 {
		fmt.Println("引数が間違っています。3つ以上のファイルは読み込めません。")
		// プロセス終了
		os.Exit(1)
	}

	// ファイルをOpenする
	logFile, err := os.Open(args[1])
	// 読み取り時の例外処理
	if err != nil {
		fmt.Println("うまくファイルを読み取れませんでした", err)
		os.Exit(1)
	}

	defer logFile.Close()

	scanner := bufio.NewScanner(logFile)
	// DB接続情報
	connectInfo := "user=test-user password=test-pass dbname=users sslmode=disable"

	db, err := sql.Open("postgres", connectInfo)
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	fmt.Println("データベース接続成功")

	// %+v を使用するとフィールド名も表示される
	// DB操作
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("error: %v", err)
	}
	
	defer tx.Rollback()
		
	for scanner.Scan() {
		line := scanner.Text()
		// テキスト一行ごとの処理		
		var logData LogData
		if err := json.Unmarshal([]byte(line), &logData); 
		err != nil {
			log.Printf("JSONパースエラー: %v", err)
			continue
		}
		insertData(ctx,db, logData, tx)
	}

	defer db.Close()
	fmt.Println("データベース接続を終了")

}