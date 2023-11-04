package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/notnil/chess"
)

const version = "0.0.1"

var dbg = false

func evaluate(move *chess.Move) float64 {
	// TODO: for now just return a random evaluation
	return (rand.Float64() - 0.5) * 400
}

func find_move(game *chess.Game) (*chess.Move, float64) {
	max_eval := math.Inf(-1)
	index_best := 0

	moves := game.ValidMoves()
	for index, move := range moves {
		evaluation := evaluate(move)
		if evaluation > max_eval {
			max_eval = evaluation
			index_best = index
		}
	}
	return moves[index_best], max_eval
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
		move, eval := find_move(game)
		fmt.Printf("info score cp %d\n", int(eval))
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
