package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/notnil/chess"
)

const version = "0.0.1"

func handle(line string) {
	tokens := strings.Split(line, " ")

	switch tokens[0] {
	case "uci":
		fmt.Println("id name Budget Snorkel")
		fmt.Println("id author Scott Lewis")
		fmt.Println("uciok")
	case "isready":
		fmt.Println("readyok")
	case "position":
		game := chess.NewGame(chess.UseNotation(chess.UCINotation{}))
		game.MoveStr("e2e4")
		game.MoveStr("e7e5")
		fmt.Println(game.Position().Board().Draw())
	case "quit":
		os.Exit(0)
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
