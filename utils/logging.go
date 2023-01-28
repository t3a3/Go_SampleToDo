package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSetting(LogFile string) {
	// os.OpenFile関数。指定したフラグで指定したファイルを開く/作成することができる。
	// fileが取得できるのでfile経由で書き込み/読み込みを行う
	// ほとんどの場合OpenやCreateの方を使う
	//os.O_RDWR=読み書き os.O_CREATE=作成 os.O_APPEND=追加
	logfile, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//go言語のlog.Fatal、log.Fatalfおよびlog.Fatallnはメッセージ出力後にos.Exit(1)を発行し、プロセスを終了しようとします。
		log.Fatalln(err)
	}
	//ログの書き込み先を標準とログファイルに
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//ログのフォーマットを指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//ログの設定先を指定
	log.SetOutput(multiLogFile)
}
