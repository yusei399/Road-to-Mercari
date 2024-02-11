// main.go
package main

import (
	"fmt"
	"os"
	"test/ex01/imgconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main <directory>")
		return
	}
	dir := os.Args[1]
	converter := imgconv.NewConverter()
	if err := converter.ConvertJPGToPNG(dir); err != nil {
		fmt.Printf("Failed to convert: %v\n", err)
		return
	}
	fmt.Println("Conversion completed.")
}

