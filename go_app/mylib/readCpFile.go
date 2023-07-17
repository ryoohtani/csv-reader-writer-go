package mylib

import (
	"fmt"
	"os"
	//strings(シャドウウィング対策)
	//"strings"
)

func WriteTextFile(newOutFile, newOutData string) error {
	outFile, err := os.Create(newOutFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = fmt.Fprintln(outFile, newOutData)
	if err != nil {
		return err
	}

	return nil
}

func CopyTextFile(csvFile, newData string) {
	existingData := ReadTextFile(csvFile)

	if existingData != "" {
		//メリット:直感的な文字列結合、デメリット:シャドウウィング
		newData = existingData + "," + newData

		//メリット:シャドウウィング対策、デメリット:stringsパッケージに依存するのがデメリット
		//newData = strings.Join([]string{existingData, newData}, ",")
	}
	outputCsv := WriteTextFile(csvFile, newData)
	
	if outputCsv != nil {
		fmt.Println("CSVファイルへの書き出しに失敗しました:", outputCsv)
		return
	}
	fmt.Println(csvFile + "に書き出しました")
}
