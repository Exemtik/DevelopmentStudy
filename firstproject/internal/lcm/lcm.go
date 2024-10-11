package lcm

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type LCMGame struct{}

func (g *LCMGame) Play(name string) {
	fmt.Printf("Welcome to the Geometric Progression Game, %s!\n", name)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	const rounds = 3
	for i := 0; i < rounds; i++ {
		nums := generateNumbers(r)
		fmt.Printf("Round %d: Find the least common multiple of: %d, %d, %d\n", i+1, nums[0], nums[1], nums[2])

		userInput := bufio.NewReader(os.Stdin)
		text, err := userInput.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		answer, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			fmt.Println("Please enter a valid answer.")
			return
		}

		correctAnswer := FindLCM(nums[0], nums[1], nums[2])

		if answer != correctAnswer {
			fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'.\n\n", answer, correctAnswer)
			return
		} else {
			fmt.Println("Correct!")
		}
	}
	// Congratulate if all answers are correct
	fmt.Printf("Congratulations, %s! You've successfully completed the LCM game.\n", name)
}

func generateNumbers(r *rand.Rand) []int {
	return []int{
		r.Intn(10) + 1,
		r.Intn(10) + 1,
		r.Intn(10) + 1,
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func FindLCM(a, b, c int) int {
	lcmAB := a * b / gcd(a, b)
	return lcmAB * c / gcd(lcmAB, c)
}
