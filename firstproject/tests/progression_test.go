package tests

import (
	"github.com/Exemtik/DevelopmentStudy/firstproject/internal/progression"
	"math/rand"
	"testing"
	"time"
)

func TestGenerateProgression(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prog := progression.GenerateProgression(r)
	if len(prog) < 5 || len(prog) > 10 {
		t.Errorf("Expected progression length between 5 and 10, got %d", len(prog))
	}
}

func TestShowProgression(t *testing.T) {
	prog := []int{1, 2, 4, 8, 16, 32, 64}
	missingIndex := 4
	output := progression.ShowProgression(prog, missingIndex)
	expected := "1 2 4 8 .. 32 64 "

	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}
