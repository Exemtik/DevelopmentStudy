package main

import (
	"github.com/Exemtik/DevelopmentStudy/firstproject/internal/app"
	"github.com/Exemtik/DevelopmentStudy/firstproject/internal/lcm"
	"github.com/Exemtik/DevelopmentStudy/firstproject/internal/progression"
)

func main() {
	runner := app.App{
		Games: map[int]app.Game{
			1: &progression.ProgressionGame{},
			2: &lcm.LCMGame{},
		},
	}
	runner.Run()
}
