package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	countLines := flag.Bool("l", false, "Count lines")
	countBytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()

	if *countLines && *countBytes {
		fmt.Printf("cannot count lines and bytes at the same time")
		os.Exit(1)
	}

	fmt.Println(Count(os.Stdin, *countLines, *countBytes))
}

func Count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if countBytes {
		for scanner.Scan() {
			return len(scanner.Bytes())
		}
	}

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
