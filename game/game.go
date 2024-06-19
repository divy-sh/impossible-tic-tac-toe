package game

import "errors"

type Game struct {
	board    [][]int
	size     int
	player   int
	moveList []Move
	winner   int
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
		board:    newBoard,
		player:   1,
		moveList: []Move{},
		winner:   0,
	}
	return game
}

func (g *Game) LegalMoves() []Move {
	moves := []Move{}
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			if g.board[i][j] == 0 {
				move := Move{
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

func (g *Game) Move(x int, y int) (Game, error) {
	move := Move{s1: x, s2: y, player: g.player}
	return g.PushMove(move)
}

func (g *Game) PushMove(move Move) (Game, error) {
	if move.s1 >= g.size || move.s2 >= g.size || g.board[move.s1][move.s2] != 0 || g.winner != 0 {
		return *g, errors.New("invalid move")
	}
	newGame := *g
	newGame.board = make([][]int, len(g.board))
	for i := range g.board {
		newGame.board[i] = make([]int, len(g.board[i]))
		copy(newGame.board[i], g.board[i])
	}
	newGame.board[move.s1][move.s2] = newGame.player
	newGame.moveList = append(newGame.moveList, move)
	newGame.updateGameStatus()
	return newGame, nil
}

func (g *Game) PrintGameStatus() string {
	if !g.IsGameOver() {
		return "game not finished"
	} else if g.winner == 1 {
		return "X gon give it to ya"
	} else if g.winner == -1 {
		return "O-nly I can win"
	} else {
		return "it's a draw... zzz"
	}
}

func (g *Game) GetGameStatus() int {
	return g.winner
}

func (g *Game) IsGameOver() bool {
	return g.winner != 0 || len(g.moveList) == g.size*g.size
}

func (g *Game) PrintBoard() string {
	boardString := ""
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			switch g.board[i][j] {
			case 0:
				boardString += "   "
			case 1:
				boardString += " X "
			case -1:
				boardString += " O "
			}
			if j < g.size-1 {
				boardString += "|"
			}
		}
		boardString += "\n"
		if i < g.size-1 {
			boardString += "---|---|---\n"
		}
	}
	return boardString
}

func (g *Game) updateGameStatus() {
	g.changePlayer()

	diagSum1 := 0
	diagSum2 := 0
	for i := 0; i < g.size; i++ {
		rowSum := 0
		colSum := 0
		for j := 0; j < g.size; j++ {
			rowSum += g.board[i][j]
			colSum += g.board[j][i]
			if rowSum == g.size || colSum == g.size {
				g.winner = 1
				break
			}
			if rowSum == -g.size || colSum == -g.size {
				g.winner = -1
				break
			}
		}
		diagSum1 += g.board[i][i]
		diagSum2 += g.board[g.size-i-1][i]
		if diagSum1 == g.size || diagSum2 == g.size {
			g.winner = 1
			break
		}
		if diagSum1 == -g.size || diagSum2 == -g.size {
			g.winner = -1
			break
		}
	}
}

func (g *Game) changePlayer() {
	if g.player == 1 {
		g.player = -1
	} else {
		g.player = 1
	}
}
