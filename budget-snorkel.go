package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/notnil/chess"
)

const version = "0.0.1"

var dbg = false

func find_move(game *chess.Game) *chess.Move {
	moves := game.ValidMoves()
	index := rand.Intn(len(moves))
	return moves[index]
}

func handle(line string, game *chess.Game) *chess.Game {
	tokens := strings.Split(line, " ")
	cmd := tokens[0]

	switch cmd {
	case "uci":
		fmt.Println("id name Budget Snorkel")
		fmt.Println("id author Scott Lewis")
		fmt.Println("uciok")

	case "isready":
		fmt.Println("readyok")

	case "stop":

	case "quit":
		os.Exit(0)

	case "debug":
		if len(tokens) == 1 {
			dbg = !dbg
		} else {
			switch cmd {
			case "on":
				dbg = true
			case "off":
				dbg = false
			default:
				dbg = !dbg
			}
		}
		fmt.Printf("debug %t\n", dbg)

	case "position":
		if len(tokens) == 1 {
			return game
		}

		if tokens[1] == "startpos" {
			game = chess.NewGame(chess.UseNotation(chess.UCINotation{}))
		}

		if len(tokens) == 2 {
			return game
		}

		for _, move := range tokens[2:] {
			game.MoveStr(move)
			// fmt.Println(game.Position().Board().Draw())
		}

	case "show":
		fmt.Println(game.Position().Board().Draw())

	case "go":
		move := find_move(game)
		fmt.Printf("info score cp %d\n", 100)
		fmt.Printf("bestmove %s\n", move.String())

	default:
		// ignore
	}
	return game
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Budget Snorkel v%s\n", version)
	game := chess.NewGame(chess.UseNotation(chess.UCINotation{}))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		game = handle(scanner.Text(), game)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
