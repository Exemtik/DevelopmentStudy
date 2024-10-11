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
	Play(name string)
}

type App struct {
	Games map[int]Game
}

func (app *App) Run() {
	name := greetUser()

	for {
		fmt.Println("Welcome to the Brain Games!")
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
				game.Play(name)
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

func greetUser() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("May I have your name? ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s!\n", name)
	return name
}
