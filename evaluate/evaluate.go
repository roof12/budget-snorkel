package evaluate

import (
	"github.com/notnil/chess"
)

var pieceValues = map[chess.PieceType]float64{
	chess.NoPieceType: 0,
	chess.King:        1000,
	chess.Queen:       9,
	chess.Rook:        5,
	chess.Bishop:      3.2,
	chess.Knight:      3,
	chess.Pawn:        1,
}

func Evaluate(game chess.Game, move *chess.Move) int16 {
	return int16((evaluatePosition(game, move) +
		evaluateOpeningDevelopment(game, move)) * 100)
}

func evaluateOpeningDevelopment(game chess.Game, move *chess.Move) float64 {
	if len(game.Moves()) > 30 {
		return 0.0
	}

	game.Move(move)

	total := 0.0
	whiteSquares := []chess.Square{
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
			total += 0.05
		}
	}
	for _, square := range blackSquares {
		if squareMap[square].Type() == chess.NoPieceType {
			total -= 0.05
		}
	}
	return total
}

func evaluatePosition(game chess.Game, move *chess.Move) float64 {
	total := 0.0
	game.Move(move)
	for _, piece := range game.Position().Board().SquareMap() {
		if piece.Color() == chess.White {
			total += pieceValues[piece.Type()]
		} else {
			total -= pieceValues[piece.Type()]
		}
	}
	return total
}
