package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game interface {
	Play()
}

type App struct {
	Games map[int]Game
}

func (app *App) Run() {
	for {
		fmt.Println("Welcome to Brain Game!")
		fmt.Println("1. Play Geometric Progression")
		fmt.Println("2. Play LCM Game")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		userInput := bufio.NewReader(os.Stdin)
		text, err := userInput.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		choice, err := strconv.Atoi(strings.TrimSpace(text))
		if err == nil {
			if game, ok := app.Games[choice]; ok {
				game.Play()
			} else if choice == 3 {
				fmt.Println("Exiting the game. Goodbye!")
				break
			} else {
				fmt.Println("Invalid choice, please select 1, 2, or 3.")
			}
		} else {
			fmt.Println("Invalid choice, please select 1, 2, or 3.")
		}
	}
}
