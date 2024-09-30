package progression

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func PlayProgressionGame() {
	fmt.Println("Welcome to the Geometric Progression Game!")

	name := getUserName()

	rand.Seed(time.Now().UnixNano())
	progress := generateProgression()
	missingIndex := rand.Intn(len(progress))

	fmt.Println("What number is missing in the progression?")
	showProgression(progress, missingIndex)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your answer: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	userAnswer, err := strconv.Atoi(userInput)
	if err != nil || userAnswer != progress[missingIndex] {
		fmt.Printf("'%s' is wrong answer ;(. Correct answer was '%d'. Let's try again, %s!\n", userInput, progress[missingIndex], name)
	} else {
		fmt.Printf("Congratulations, %s!\n", name)
	}
}

func generateProgression() []int {
	length := rand.Intn(5) + 5
	base := rand.Intn(5) + 2
	start := rand.Intn(10) + 1

	progression := make([]int, length)
	progression[0] = start
	for i := 1; i < length; i++ {
		progression[i] = progression[i-1] * base
	}

	return progression
}

func showProgression(prog []int, missingIndex int) {
	for i, val := range prog {
		if i == missingIndex {
			fmt.Print(".. ")
		} else {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()
}

func getUserName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Can you write your name? ")
	name, _ := reader.ReadString('\n')
	return strings.TrimSpace(name)
}
