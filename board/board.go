package board

import (
	"fmt"
	tm "github.com/buger/goterm"
	"math/rand"
)

// Board tiles are numbered as follows:
//  0  1  2  3
//  4  5  6  7
//  8  9 10 11
// 12 13 14 15
type board struct {
	board [4][4]int
	score int
}

func New() *board {
	// initialize random seed based on current time
	b := &board{}
	b.score = 0
	for i := 0; i < 3; i++ {
		b.populateTile(b.getRandomTile())
	}
	return b
}

func (b *board) Print() {
	// Clear the terminal screen
	tm.Clear()
	tm.Flush()
	tm.MoveCursor(1, 1)

	fmt.Println("\nScore: ", b.score)
	fmt.Println("-----------------------------")
	for i := 0; i < 4; i++ {
		fmt.Println("|", formatTile(b.board[i][0]), "|", formatTile(b.board[i][1]), "|",
			formatTile(b.board[i][2]), "|", formatTile(b.board[i][3]), "|")
		if i == 3 {
		}
	}
	fmt.Println("-----------------------------")
}
func (b *board) IsGameOver() bool {
	if b.getAllFreeTiles() == nil {
		for i := 0; i < 4; i++ {
			if !isStuck(b.board[i]) {
				return false
			}
			column := [4]int{b.board[0][i], b.board[1][i], b.board[2][i], b.board[3][i]}
			if !isStuck(column) {
				return false
			}
		}
		return true
	}
	return false
}

func (b *board) ProcessMove(input string) {
	oldBoard := b.board
	switch input {
	case "w":
		b.sumUp()
		break
	case "a":
		b.sumLeft()
		break
	case "s":
		b.sumDown()
		break
	case "d":
		b.sumRight()
		break
	}
	if !b.isEqualTo(oldBoard) {
		b.spawnTile()
	}
}

func (b *board) spawnTile() {
	b.populateTile(b.getRandomTile())
}

func (b *board) sumLeft() {
	for i := 0; i < 4; i++ {
		if !isStuck(b.board[i]) {
			b.board[i] = b.processRowLeft(b.board[i])
		}
	}
}

func (b *board) sumRight() {
	for i := 0; i < 4; i++ {
		row := [4]int{b.board[i][3], b.board[i][2], b.board[i][1], b.board[i][0]}
		if !isStuck(row) {
			row = b.processRowLeft(row)
			for j := 0; j < 4; j++ {
				b.board[i][3-j] = row[j]
			}
		}
	}
}

func (b *board) sumUp() {
	for i := 0; i < 4; i++ {
		column := [4]int{b.board[0][i], b.board[1][i], b.board[2][i], b.board[3][i]}
		if !isStuck(column) {
			column = b.processRowLeft(column)
			for j := 0; j < 4; j++ {
				b.board[j][i] = column[j]
			}
		}
	}
}

func (b *board) sumDown() {
	for i := 0; i < 4; i++ {
		column := [4]int{b.board[3][i], b.board[2][i], b.board[1][i], b.board[0][i]}
		if !isStuck(column) {
			column = b.processRowLeft(column)
			for j := 0; j < 4; j++ {
				b.board[3-j][i] = column[j]
			}
		}
	}
}

func (b *board) isEqualTo(other [4][4]int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.board[i][j] != other[i][j] {
				return false
			}
		}
	}
	return true
}

func formatTile(i int) string {
	if i == 0 {
		return "    "
	}
	return fmt.Sprintf("%4d", i)
}

func (b *board) getAllFreeTiles() []int {
	var tiles []int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.board[i][j] == 0 {
				tiles = append(tiles, i*4+j)
			}
		}
	}
	return tiles
}

func (b *board) getRandomTile() int {
	tiles := b.getAllFreeTiles()
	return tiles[rand.Intn(len(tiles))]
}

func (b *board) populateTile(tileNumber int) {
	i := tileNumber / 4
	j := tileNumber % 4
	b.board[i][j] = 2 + (rand.Intn(2) * 2)
}

func isStuck(row [4]int) bool {
	for i := 0; i < 4; i++ {
		if row[i] == 0 {
			return false
		}
	}
	return !(row[0] == row[1] || row[1] == row[2] || row[2] == row[3])
}

func (b *board) processRowLeft(row [4]int) [4]int {
	// Compress excess zeros
	noExcessZeros := false
	for !noExcessZeros {
		for i := 0; i < 3; i++ {
			if row[i] == 0 {
				row[i] = row[i+1]
				row[i+1] = 0
				if i == 2 {
					row[3] = 0
				}
			}
		}
		noExcessZeros = row[0] == 0 && row[1] == 0 && row[2] == 0 && row[3] == 0 ||
			row[0] != 0 && row[1] == 0 && row[2] == 0 && row[3] == 0 ||
			row[0] != 0 && row[1] != 0 && row[2] == 0 && row[3] == 0 ||
			row[0] != 0 && row[1] != 0 && row[2] != 0
	}

	// sum adjacent tiles
	for j := 0; j < 3; j++ {
		if row[j] == row[j+1] {
			row[j] += row[j+1]
			b.score += row[j]
			row[j+1] = 0
		}
	}

	// Shift to the left
	for j := 0; j < 3; j++ {
		if row[j] == 0 {
			row[j] = row[j+1]
			row[j+1] = 0
		}
	}
	return row
}
