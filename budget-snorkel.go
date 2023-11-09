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

const version = "0.0.2"

var pieceValues = map[chess.PieceType]float64{
	chess.NoPieceType: 0,
	chess.King:        1000,
	chess.Queen:       9,
	chess.Rook:        5,
	chess.Bishop:      3.2,
	chess.Knight:      3,
	chess.Pawn:        1,
}

var dbg = false

func evaluate(game chess.Game, move *chess.Move) int16 {
	total := 0.0
	game.Move(move)
	for _, piece := range game.Position().Board().SquareMap() {
		if piece.Color() == chess.White {
			total += pieceValues[piece.Type()]
		} else {
			total -= pieceValues[piece.Type()]
		}
	}
	return int16(total * 100)
}

func findMove(game chess.Game) (*chess.Move, int16) {
	blackToMove := game.Position().Turn() == chess.Black
	var maxEval int16 = math.MinInt16
	bestMoves := []*chess.Move{}
	moves := game.ValidMoves()
	for _, move := range moves {
		evaluation := evaluate(game, move)
		if blackToMove {
			evaluation *= -1
		}

		if evaluation > maxEval {
			maxEval = evaluation
			bestMoves = []*chess.Move{move}
		} else if evaluation == maxEval {
			bestMoves = append(bestMoves, move)
		}
	}
	return bestMoves[rand.Intn(len(bestMoves))], maxEval
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

		if len(tokens) <= 3 {
			return game
		}

		for _, move := range tokens[3:] {
			game.MoveStr(move)
			// fmt.Println(game.Position().Board().Draw())
		}

	case "show":
		fmt.Println(game.Position().Board().Draw())

	case "go":
		move, eval := findMove(*game)
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
