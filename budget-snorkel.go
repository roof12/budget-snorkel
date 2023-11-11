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

func findMove(game chess.Game, depth int8) (*chess.Move, int16) {
	moves := game.ValidMoves()
	if len(moves) == 0 {
		return nil, 0
	}

	blackToMove := game.Position().Turn() == chess.Black
	var bestEval int16 = math.MinInt16
	bestMoves := []*chess.Move{}
	for _, move := range moves {
		evaluation := evaluate(game, move)

		if depth > 0 {
			newGame := game.Clone()
			newGame.Move(move)
			newMove, newEval := findMove(*newGame, depth-1)
			if newMove != nil {
				evaluation = newEval
			}
		}

		if (len(bestMoves) == 0) ||
			(blackToMove && evaluation < bestEval) ||
			(!blackToMove && evaluation > bestEval) {
			bestEval = evaluation
			bestMoves = []*chess.Move{move}
		} else if evaluation == bestEval {
			bestMoves = append(bestMoves, move)
		}
	}

	randInt := rand.Intn(len(bestMoves))
	return bestMoves[randInt], bestEval
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
		move, eval := findMove(*game, 1)
		fmt.Printf("info score cp %d\n", int(eval))
		if move == nil {
			fmt.Println("bestmove (none)")
		} else {
			fmt.Printf("bestmove %s\n", move.String())
		}

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
