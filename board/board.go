package board

import (
	"fmt"
	"math/rand"
)

// Board tiles are numbered as follows:
//  0  1  2  3
//  4  5  6  7
//  8  9 10 11
// 12 13 14 15
type board struct {
	board [4][4]int
}

func New() *board {
	// initialize random seed based on current time
	b := &board{}
	for i := 0; i < 7; i++ {
		b.populateTile(b.getRandomTile())
	}
	return b
}

func (b *board) Print() {
	for i := 0; i < 4; i++ {
		if i == 0 {
			fmt.Println("-----------------")
		}
		fmt.Println("|", formatTile(b.board[i][0]), "|", formatTile(b.board[i][1]), "|",
			formatTile(b.board[i][2]), "|", formatTile(b.board[i][3]), "|")
		if i == 3 {
			fmt.Println("-----------------")
		}
	}
}

func formatTile(i int) string {
	if i == 0 {
		return " "
	}
	return fmt.Sprintf("%d", i)
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

func (b *board) IsGameOver() bool {
	return b.getAllFreeTiles() == nil
}

func (b *board) SumLeft() {
	for i := 0; i < 4; i++ {
		if !isStuck(b.board[i]) {
			b.board[i] = processRowLeft(b.board[i])
		}
	}
}

func (b *board) SumRight() {
	for i := 0; i < 4; i++ {
		if !isStuck(b.board[i]) {
			// Reverse the row
			row := [4]int{b.board[i][3], b.board[i][2], b.board[i][1], b.board[i][0]}
			row = processRowLeft(row)
			// Reverse the row again
			for j, k := 0, len(row)-1; j < k; j, k = j+1, k-1 {
				row[j], row[k] = row[k], row[j]
			}
			b.board[i] = row
		}
	}
}

func isStuck(row [4]int) bool {
	for j := 0; j < 4; j++ {
		if row[j] != 0 {
			return false
		}
	}
	return row[0] == row[1] || row[1] == row[2] || row[2] == row[3]
}

func processRowLeft(row [4]int) [4]int {
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

	// Sum adjacent tiles
	for j := 0; j < 3; j++ {
		if row[j] == row[j+1] {
			row[j] += row[j+1]
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
