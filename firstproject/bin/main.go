package main

import (
	"firstproject/src/lcm"
	"firstproject/src/progression"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Welcome to Brain Game!")
		fmt.Println("1. Play Geometric Progression")
		fmt.Println("2. Play NOK")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			progression.PlayProgressionGame()
		case 2:
			lcm.PlayLCMGame()
		case 3:
			fmt.Println("Exiting the game. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please select 1, 2, or 3.")
		}
	}
}
