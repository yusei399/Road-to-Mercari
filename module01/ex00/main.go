package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ft_cat(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		_, err := fmt.Fprintln(writer, scanner.Text())
		if err != nil {
			return err
		}
	}
	return scanner.Err() 
}

func main() {
	if len(os.Args) < 2 {
		ft_cat(os.Stdin, os.Stdout)
	}

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr,  err)
			continue 
		}
		err = ft_cat(file, os.Stdout)
		if err != nil {
			fmt.Fprintln(os.Stderr,  err)
		}
		file.Close()
	}
}

