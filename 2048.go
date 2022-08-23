package main

import (
	"go2048/board"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	b := board.New()
	b.Print()
	b.SumRight()
	b.Print()
}
