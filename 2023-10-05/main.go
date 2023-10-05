package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// 今日の日付でディレクトリを作成する
	todayString := time.Now().Format("2006-01-02")
	dirString := "../" + todayString
	// MkdirAll で親ディレクトリを含めてディレクトリ階層を指定する
	if err := os.MkdirAll(dirString, 0777); err != nil {
		log.Fatal(err)
	}

	if err := os.Chdir(dirString); err != nil {
		log.Fatal(err)
	}

	// ディレクトリ内部にmain.go, go.modを作成する
	mainFile, err := os.Create("main.go")
	if err != nil {
		log.Fatal(err)
	}
	modFile, modCreateErr := os.Create("go.mod")
	if modCreateErr != nil {
		log.Fatal(err)
	}

	// 無名関数で最後にファイルを閉じる
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
	// 改行は \n で行い、ダブルクォーテーションで囲まないと改行にならない
	_, err = mainFile.WriteString("package main\n\nfunc main() {}")
	if err != nil {
		log.Fatal(err)
	}

	modFileString := "module " + todayString + "\n\ngo 1.20"
	_, err = modFile.WriteString(modFileString)
	if err != nil {
		log.Fatal(err)
	}
}
