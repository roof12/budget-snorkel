package evaluate

import (
	"fmt"

	"github.com/notnil/chess"
)

type Delta float64

const (
	Casteled  Delta = 1.0
	Developed Delta = 0.05
)

var pieceValues = map[chess.PieceType]Delta{
	chess.NoPieceType: 0,
	chess.King:        1000,
	chess.Queen:       9,
	chess.Rook:        5,
	chess.Bishop:      3.2,
	chess.Knight:      3,
	chess.Pawn:        1,
}

func Evaluate(game chess.Game, move *chess.Move) int16 {
	return (int16(100 * (evaluatePosition(game, move) +
		evaluateOpeningDevelopment(game, move) +
		evaluateCastling(game, move))))
}

func evaluateCastling(game chess.Game, move *chess.Move) Delta {
	castleCount := 0
	var total Delta = 0.0
	game.Move(move)

	for _, mh := range game.MoveHistory() {
		if mh.Move.HasTag(chess.KingSideCastle) || mh.Move.HasTag(chess.QueenSideCastle) {
			if mh.PrePosition.Turn() == chess.White {
				total += Casteled
				fmt.Println("White has castled")
			} else {
				total -= Casteled
				fmt.Println("Black has castled")
			}
			castleCount += 1
			if castleCount == 2 {
				break
			}
		}
	}
	fmt.Println("evaluateCastling", total)
	return total
}

func evaluateOpeningDevelopment(game chess.Game, move *chess.Move) Delta {
	if len(game.Moves()) > 30 {
		fmt.Println("opening move ct", len(game.Moves()))
		return 0.0
	}

	game.Move(move)

	var total Delta = 0.0
	whiteSquares := []chess.Square{
		chess.D2,
		chess.E2,
		chess.A1,
		chess.B1,
		chess.C1,
		chess.D1,
		chess.E1,
		chess.F1,
		chess.G1,
		chess.H1,
	}
	blackSquares := []chess.Square{
		chess.D7,
		chess.E7,
		chess.A8,
		chess.B8,
		chess.C8,
		chess.D8,
		chess.E8,
		chess.F8,
		chess.G8,
		chess.H8,
	}

	squareMap := game.Position().Board().SquareMap()
	for _, square := range whiteSquares {
		if squareMap[square].Type() == chess.NoPieceType {
			fmt.Println("white developed", square)
			total += Developed
		}
	}
	for _, square := range blackSquares {
		if squareMap[square].Type() == chess.NoPieceType {
			fmt.Println("black developed", square)
			total -= Developed
		}
	}
	fmt.Println("evaluateOpeningDevelopment", total)
	return total
}

func evaluatePosition(game chess.Game, move *chess.Move) Delta {
	var total Delta = 0.0
	game.Move(move)
	for _, piece := range game.Position().Board().SquareMap() {
		if piece.Color() == chess.White {
			total += pieceValues[piece.Type()]
		} else {
			total -= pieceValues[piece.Type()]
		}
	}
	fmt.Println("evaluatePosition", total)
	return total
}
