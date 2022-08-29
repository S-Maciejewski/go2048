package main

import (
	"bufio"
	"fmt"
	tm "github.com/buger/goterm"
	"go2048/board"
	"golang.org/x/term"
	"math/rand"
	"os"
	"strings"
	"time"
)

// TODO: Use a better display method
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

	tm.Clear()
	tm.Flush()
	tm.MoveCursor(1, 1)

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
		b.ProcessMove(string(input))
		b.Print()
	}
	fmt.Println("Game over!")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
