package game

import "fmt"

type Game struct {
	board    [][]int
	size     int
	player   int
	moveList []Move
	winner   int
}

func NewGame(size int) *Game {
	newBoard := [][]int{}
	for i := 0; i < size; i++ {
		newRow := []int{}
		for j := 0; j < size; j++ {
			newRow = append(newRow, -1)
		}
		newBoard = append(newBoard, newRow)
	}
	game := &Game{
		size:     size,
		board:    newBoard,
		player:   0,
		moveList: []Move{},
	}
	return game
}

func (g *Game) LegalMoves() []*Move {
	moves := []*Move{}
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			if g.board[i][j] == -1 {
				move := &Move{
					s1:     i,
					s2:     j,
					player: g.player,
				}
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func (g *Game) Move(x int, y int) bool {
	if x >= g.size || y >= g.size || g.board[x][y] != -1 || g.winner != -1 {
		return false
	}
	g.board[x][y] = g.player
	g.moveList = append(g.moveList, Move{s1: x, s2: y, player: g.player})
	g.changePlayer()
	return true
}

func (g *Game) PushMove(move Move) bool {
	if move.s1 >= g.size || move.s2 >= g.size || g.board[move.s1][move.s1] != -1 || g.winner != -1 {
		return false
	}
	g.board[move.s1][move.s2] = g.player
	g.moveList = append(g.moveList, move)
	g.changePlayer()
	return true
}

func (g *Game) Revert() bool {
	if len(g.moveList) <= 0 {
		return false
	}
	revertMove := g.moveList[len(g.moveList)-1]
	g.board[revertMove.s1][revertMove.s2] = -1
	g.moveList = g.moveList[:len(g.moveList)-1]
	g.changePlayer()
	return true
}

func (g *Game) IsGameOver() bool {
	return g.winner != -1
}

func (g *Game) PrintBoard() {
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			switch g.board[i][j] {
			case 0:
				fmt.Print(" O ")
			case 1:
				fmt.Print(" X ")
			case -1:
				fmt.Print("   ")
			}
			if j < g.size-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < g.size-1 {
			fmt.Println("---|---|---")
		}
	}
}

func (g *Game) changePlayer() {
	if g.player == 0 {
		g.player = 1
	} else {
		g.player = 0
	}
}

func (g *Game) CheckGameStatus() bool {
	if g.winner != -1 {
		return true
	}
	if len(g.moveList) < g.size*g.size {
		return false
	}
	return true
}
