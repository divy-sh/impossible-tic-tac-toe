package main

import (
	"tic-tac-toe/game"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	board      game.Game
	buttons    [][]*widget.Button
	statusText *widget.Label
)

func main() {
	a := app.New()
	w := a.NewWindow("Tic Tac Toe")

	board = game.NewGame(3)
	buttons = make([][]*widget.Button, 3)
	grid := container.NewGridWithColumns(3)

	for i := range buttons {
		buttons[i] = make([]*widget.Button, 3)
		for j := range buttons[i] {
			i, j := i, j // Capture i, j for each button
			buttons[i][j] = widget.NewButton(" ", func() {
				buttonClicked(i, j)
			})
			grid.Add(buttons[i][j])
		}
	}

	statusText = widget.NewLabel("Player X's turn")
	content := container.NewVBox(grid, statusText)

	w.SetContent(content)
	w.ShowAndRun()
}

func buttonClicked(x, y int) {
	if board.IsGameOver() {
		return
	}

	if board.Board[x][y] != 0 {
		statusText.SetText("Invalid move, try again")
		return
	}

	newBoard, err := board.Move(x, y)
	if err != nil {
		statusText.SetText("Invalid move, try again")
		return
	}

	board = newBoard
	updateUI()

	if board.IsGameOver() {
		statusText.SetText(board.PrintGameStatus())
		return
	}

	move := Eval(board)
	board, _ = board.PushMove(*move)
	updateUI()

	if board.IsGameOver() {
		statusText.SetText(board.PrintGameStatus())
		return
	}
}

func updateUI() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch board.Board[i][j] {
			case 1:
				buttons[i][j].SetText("X")
				buttons[i][j].Disable()
			case -1:
				buttons[i][j].SetText("O")
				buttons[i][j].Disable()
			case 0:
				buttons[i][j].SetText(" ")
				buttons[i][j].Enable()
			}
		}
	}

	if board.Player == 1 {
		statusText.SetText("Player X's turn")
	} else {
		statusText.SetText("Player O's turn")
	}
}

// package main

// import (
// 	"fmt"
// 	"tic-tac-toe/game"
// )

// func main() {
// 	board := game.NewGame(3)
// 	for !board.IsGameOver() {
// 		fmt.Println(board.PrintBoard())
// 		var x int
// 		fmt.Printf("move: ")
// 		fmt.Scanf("%d", &x)
// 		if 0 < x && x < 10 {
// 			newBoard, err := board.Move((x-1)/3, (x-1)%3)
// 			if err != nil {
// 				fmt.Println("invalid input")
// 				continue
// 			} else {
// 				board = newBoard
// 			}
// 		} else {
// 			fmt.Println("invalid input, try again")
// 			continue
// 		}
// 		if board.IsGameOver() {
// 			break
// 		}
// 		move := Eval(board)
// 		board, _ = board.PushMove(*move)
// 	}
// 	fmt.Println(board.PrintBoard())
// 	fmt.Println(board.PrintGameStatus())
// }
