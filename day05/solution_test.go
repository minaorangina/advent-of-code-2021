package day05

import (
	_ "embed"
	"fmt"
	"testing"
)

var (
	//go:embed testinput.txt
	testInput []byte
	//go:embed input.txt
	input []byte
)

func TestPart1(t *testing.T) {
	got := Part1(testInput)

	if got != 5 {
		t.Fatalf("got %d, want 5", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part1(input))
	})
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)

	if got != 12 {
		t.Fatalf("got %d, want 12", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part2(input))
	})
}
