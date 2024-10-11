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

type ProgressionGame struct{}

func (g *ProgressionGame) Play(name string) {
	fmt.Printf("Welcome to the Geometric Progression Game, %s!\n", name)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	const rounds = 3
	for i := 0; i < rounds; i++ {
		progress := GenerateProgression(r)
		missingIndex := r.Intn(len(progress))

		fmt.Printf("Round %d: What number is missing in the progression?\n", i+1)
		fmt.Println(ShowProgression(progress, missingIndex))

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Your answer: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		userAnswer, err := strconv.Atoi(userInput)
		if err != nil || userAnswer != progress[missingIndex] {
			fmt.Printf("'%s' is wrong answer ;(. Correct answer was '%d'. Let's try again, %s!\n\n", userInput, progress[missingIndex], name)
			return
		} else {
			fmt.Println("Correct!")
		}
	}
	fmt.Printf("Congratulations, %s! You've successfully completed the Progression game.\n", name)
}

func GenerateProgression(r *rand.Rand) []int {
	length := r.Intn(5) + 5
	base := r.Intn(5) + 2
	start := r.Intn(10) + 1

	progression := make([]int, length)
	progression[0] = start
	for i := 1; i < length; i++ {
		progression[i] = progression[i-1] * base
	}

	return progression
}

func ShowProgression(prog []int, missingIndex int) string {
	result := ""
	for i, val := range prog {
		if i == missingIndex {
			result += ".. "
		} else {
			result += fmt.Sprintf("%d ", val)
		}
	}
	return result
}
