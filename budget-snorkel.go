package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const version = "0.0.1"

func handle(line string) {
	tokens := strings.Split(line, " ")

	switch tokens[0] {
	case "uci":
		fmt.Println("id name Budget Snorkel")
		fmt.Println("id author Scott Lewis")
		fmt.Println("uciok")
	default:
	}
}

func main() {
	fmt.Printf("Budget Snorkel v%s\n", version)
	scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
		handle(scanner.Text())
	}
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
