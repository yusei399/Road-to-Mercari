package main

import (
	"fmt"
	"os"
	"convert/converter" // 自作パッケージのインポート
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: invalid argument")
		return
	}

	dir := os.Args[1]
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("error: %s: no such file or directory\n", dir)
		return
	}

	c := converter.NewConverter()
	err := c.ConvertJPGToPNG(dir)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}

