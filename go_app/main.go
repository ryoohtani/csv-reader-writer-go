package main

import (
	"fmt"
	"go_app/mylib"
)

func main() {
	var csvFile string
	fmt.Println("csvファイルへの記載を行なってください")
	//&でcsvFile変数のポインタを取得
	//Scanln関数はエンターキーが押された時点で入力確定
	fmt.Scanln(&csvFile)

	if csvFile != "" {

		mylib.NewTextFile(csvFile)

		readResults := mylib.ReadTextFile("input.csv")
		fmt.Println("input.csvファイルの内容:", readResults)

		mylib.CopyTextFile("output.csv", readResults)

	} else {
		fmt.Println("csvファイルへの記載がされなかったのでアプリを終了します")
	}
}