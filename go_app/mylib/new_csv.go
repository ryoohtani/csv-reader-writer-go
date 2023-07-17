package mylib

import (
	"fmt"
	"os"
)

func NewTextFile(csvFile string) {
	file, err := os.Create("input.csv")

	//エラー発生(真であれば)したらcsvの作成失敗
	if err != nil {
		fmt.Println("csvファイルの作成に失敗しました:", err)
		return
	}
	defer file.Close()

	//csvFileでcsvの内容、fileはos.Createで作成されたファイル、Fprintln関数でファイルへ出力
	_, err = fmt.Fprintln(file, csvFile)
	if err != nil {
		fmt.Println("csvファイルへの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("csvファイルを出力しました")
}