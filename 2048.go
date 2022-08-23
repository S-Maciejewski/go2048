package main

import (
	"fmt"
	"go2048/board"
	"golang.org/x/term"
	"math/rand"
	"os"
	"strings"
	"time"
)

// TODO: Spawn new tiles only after a valid move; count score; use a better display method (overwrite previous board)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	// Setup reading form terminal without buffering
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {
			panic(err)
		}
	}(int(os.Stdin.Fd()), oldState)

	// Initialize board
	b := board.New()
	b.Print()

	// Game loop
	for {
		if b.IsGameOver() {
			break
		}
		fmt.Print("Enter your move using w, a, s, d keys: ")
		input := make([]byte, 1)
		for inputCorrect := false; !inputCorrect; inputCorrect = strings.Contains("wasd", string(input)) {
			_, err = os.Stdin.Read(input)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println(string(input), " read")
		switch string(input) {
		case "w":
			b.SumUp()
			break
		case "a":
			b.SumLeft()
			break
		case "s":
			b.SumDown()
			break
		case "d":
			b.SumRight()
			break
		}

		b.SpawnTile()
		b.Print()
	}
	fmt.Println("Game over!")
}
