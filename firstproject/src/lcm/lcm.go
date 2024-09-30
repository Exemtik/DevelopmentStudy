package lcm

import (
	"fmt"
	"math/rand"
	"time"
)

func PlayLCMGame() {
	fmt.Println("Welcome to the Least Common Multiple Game!")

	rand.Seed(time.Now().UnixNano())
	nums := generateNumbers()

	fmt.Printf("Find the least common multiple of: %d, %d, %d\n", nums[0], nums[1], nums[2])

	var answer int
	fmt.Scanln(&answer)

	correctAnswer := findLCM(nums[0], nums[1], nums[2])

	if answer == correctAnswer {
		fmt.Println("Correct!")
	} else {
		fmt.Printf("Wrong! The correct answer was %d.\n", correctAnswer)
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

func findLCM(a, b, c int) int {
	lcmAB := a * b / gcd(a, b)
	return lcmAB * c / gcd(lcmAB, c)
}
