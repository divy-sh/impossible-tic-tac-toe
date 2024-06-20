package game

import "errors"

type Game struct {
	Board    [][]int
	size     int
	Player   int
	moveList []Move
	Winner   int
}

func NewGame(size int) Game {
	newBoard := [][]int{}
	for i := 0; i < size; i++ {
		newRow := []int{}
		for j := 0; j < size; j++ {
			newRow = append(newRow, 0)
		}
		newBoard = append(newBoard, newRow)
	}
	game := Game{
		size:     size,
		Board:    newBoard,
		Player:   1,
		moveList: []Move{},
		Winner:   0,
	}
	return game
}

func (g *Game) LegalMoves() []Move {
	moves := []Move{}
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			if g.Board[i][j] == 0 {
				move := Move{
					s1:     i,
					s2:     j,
					player: g.Player,
				}
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func (g *Game) Move(x int, y int) (Game, error) {
	move := Move{s1: x, s2: y, player: g.Player}
	return g.PushMove(move)
}

func (g *Game) PushMove(move Move) (Game, error) {
	if move.s1 >= g.size || move.s2 >= g.size || g.Board[move.s1][move.s2] != 0 || g.Winner != 0 {
		return *g, errors.New("invalid move")
	}
	newGame := *g
	newGame.Board = make([][]int, len(g.Board))
	for i := range g.Board {
		newGame.Board[i] = make([]int, len(g.Board[i]))
		copy(newGame.Board[i], g.Board[i])
	}
	newGame.Board[move.s1][move.s2] = newGame.Player
	newGame.moveList = append(newGame.moveList, move)
	newGame.updateGameStatus()
	return newGame, nil
}

func (g *Game) PrintGameStatus() string {
	if !g.IsGameOver() {
		return "game not finished"
	} else if g.Winner == 1 {
		return "X gon give it to ya"
	} else if g.Winner == -1 {
		return "O-nly I can win"
	} else {
		return "it's a draw... zzz"
	}
}

func (g *Game) GetGameStatus() int {
	return g.Winner
}

func (g *Game) IsGameOver() bool {
	return g.Winner != 0 || len(g.moveList) == g.size*g.size
}

func (g *Game) PrintBoard() string {
	BoardString := ""
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			switch g.Board[i][j] {
			case 0:
				BoardString += "   "
			case 1:
				BoardString += " X "
			case -1:
				BoardString += " O "
			}
			if j < g.size-1 {
				BoardString += "|"
			}
		}
		BoardString += "\n"
		if i < g.size-1 {
			BoardString += "---|---|---\n"
		}
	}
	return BoardString
}

func (g *Game) updateGameStatus() {
	g.changePlayer()

	diagSum1 := 0
	diagSum2 := 0
	for i := 0; i < g.size; i++ {
		rowSum := 0
		colSum := 0
		for j := 0; j < g.size; j++ {
			rowSum += g.Board[i][j]
			colSum += g.Board[j][i]
			if rowSum == g.size || colSum == g.size {
				g.Winner = 1
				break
			}
			if rowSum == -g.size || colSum == -g.size {
				g.Winner = -1
				break
			}
		}
		diagSum1 += g.Board[i][i]
		diagSum2 += g.Board[g.size-i-1][i]
		if diagSum1 == g.size || diagSum2 == g.size {
			g.Winner = 1
			break
		}
		if diagSum1 == -g.size || diagSum2 == -g.size {
			g.Winner = -1
			break
		}
	}
}

func (g *Game) changePlayer() {
	if g.Player == 1 {
		g.Player = -1
	} else {
		g.Player = 1
	}
}
