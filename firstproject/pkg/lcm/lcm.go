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

func (g *LCMGame) Play() {
	fmt.Println("Welcome to the Least Common Multiple Game!")

	rand.Seed(time.Now().UnixNano())
	nums := generateNumbers()

	fmt.Printf("Find the least common multiple of: %d, %d, %d\n", nums[0], nums[1], nums[2])

	userInput := bufio.NewReader(os.Stdin)
	text, err := userInput.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	answer, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		fmt.Println("Please enter a valid answer.")
	} else {
		correctAnswer := FindLCM(nums[0], nums[1], nums[2])

		if answer == correctAnswer {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Wrong! The correct answer was %d.\n", correctAnswer)
		}
	}
}

func generateNumbers() []int {
	return []int{
		rand.Intn(50) + 1,
		rand.Intn(50) + 1,
		rand.Intn(50) + 1,
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
