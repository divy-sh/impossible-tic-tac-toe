package main

import (
	"fmt"
	"tic-tac-toe/game"
)

func main() {
	board := game.NewGame(3)
	for !board.IsGameOver() {
		fmt.Println(board.PrintBoard())
		var x int
		fmt.Printf("move: ")
		fmt.Scanf("%d", &x)
		if 0 < x && x < 10 {
			newBoard, err := board.Move((x-1)/3, (x-1)%3)
			if err != nil {
				fmt.Println("invalid input")
				continue
			} else {
				board = newBoard
			}
		} else {
			fmt.Println("invalid input, try again")
			continue
		}
		if board.IsGameOver() {
			break
		}
		move := Eval(board)
		board, _ = board.PushMove(*move)
	}
	fmt.Println(board.PrintBoard())
	fmt.Println(board.PrintGameStatus())
}
