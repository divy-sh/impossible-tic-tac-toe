package main

import (
	"fmt"
	"math"
	"tic-tac-toe/game"
)

func Eval(game game.Game) *game.Move {
	bestScore := -math.MaxFloat32
	moves := game.LegalMoves()
	if len(moves) == 0 {
		return nil
	}
	bestMove := moves[0]
	for _, move := range moves {
		newBoard, _ := game.PushMove(move)
		score := -negamax(newBoard, -math.MaxFloat64, math.MaxFloat64, false)
		fmt.Println(move, score)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return &bestMove
}

func negamax(game game.Game, alpha, beta float64, isMaximizingPlayer bool) float64 {
	if game.IsGameOver() {
		eval := float64(game.GetGameStatus())
		if isMaximizingPlayer {
			eval *= -1
		}
		return eval
	}
	bestScore := -math.MaxFloat64
	moves := game.LegalMoves()
	for _, move := range moves {
		newBoard, _ := game.PushMove(move)
		score := -negamax(newBoard, -beta, -alpha, !isMaximizingPlayer)
		bestScore = max(bestScore, score)
		alpha = max(alpha, score)
		if alpha >= beta {
			break
		}
	}
	return bestScore
}
