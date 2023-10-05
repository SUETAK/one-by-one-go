package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// 今日の日付でディレクトリを作成する
	//todayString := time.Now().Format("2006-01-02")
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	todayString := tomorrow.Format("2006-01-02")
	dirString := "../" + todayString
	if err := os.MkdirAll(dirString, 0777); err != nil {
		log.Fatal(err)
	}

	// ディレクトリ内部にmain.go, go.modを作成する
	if err := os.Chdir(todayString); err != nil {
		log.Fatal(err)
	}

	// main.go
	mainFile, err := os.Create("main.go")
	if err != nil {
		log.Fatal(err)
	}

	// go.mod
	modFile, modCreateErr := os.Create("go.mod")
	if modCreateErr != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := mainFile.Close(); err != nil {
			log.Fatal(err)
		}
		if err := modFile.Close(); err != nil {
			log.Fatal(err)
		}

		log.Default().Println("Done!")
	}()

	// main ファイルに書き込む
	_, err = mainFile.WriteString(`package main \n \n func main() {}`)
	if err != nil {
		log.Fatal(err)
	}

	modFileString := `module ` + todayString + `/n /n go 1.20`
	_, err = modFile.WriteString(modFileString)
	if err != nil {
		log.Fatal(err)
	}
}
