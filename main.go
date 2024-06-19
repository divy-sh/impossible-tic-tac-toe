package main

import (
	"fmt"
	"tic-tac-toe/game"
)

func main() {
	board := game.NewGame(3)
	for !board.IsGameOver() {
		board.PrintBoard()
		var x int
		fmt.Printf("move: ")
		fmt.Scanf("%d", &x)
		board.Move((x-1)/3, (x-1)%3)
	}
	board.PrintBoard()
}
