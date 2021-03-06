package day07

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/minaorangina/advent-of-code-2021/parse"
)

var (
	//go:embed testinput.txt
	testInput []byte
	//go:embed input.txt
	input []byte
)

func TestPart1(t *testing.T) {
	got := Part1((parse.ToIntSlice(testInput)))
	if got != 37 {
		t.Errorf("got %d, want 37", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part1(parse.ToIntSlice(input)))
	})
}

func TestPart2(t *testing.T) {
	got := Part2(parse.ToIntSlice(testInput))
	if got != 168 {
		t.Errorf("got %d, want 168", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part2(parse.ToIntSlice(input)))
	})
}
