package mylib

import (
	"fmt"
	"os"
)

//string { は最終的にReadTextFile関数に返す値をstringにするため
func ReadTextFile(csvFile string) string {
	readFile, err := os.ReadFile(csvFile)
	if err != nil {
		fmt.Println("csvファイルの読み込みに失敗しました:", err)
		return ""
	}
	return string(readFile)
}