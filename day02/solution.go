package day02

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)

	submarine := coords{}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(string(line))

		if len(parts) != 2 {
			return 0, fmt.Errorf("found invalid line: %v", parts)
		}
		if !validCommand(parts[0]) {
			return 0, fmt.Errorf("bad input: %s", parts[0])
		}

		var num int
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		submarine.move(parts[0], num)
	}

	return submarine.multiply(), nil
}

func Part2(input []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)

	submarine := coords{}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(string(line))

		if len(parts) != 2 {
			return 0, fmt.Errorf("found invalid line: %v", parts)
		}
		if !validCommand(parts[0]) {
			return 0, fmt.Errorf("bad input: %s", parts[0])
		}

		var num int
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		submarine.moveWithAim(parts[0], num)
	}

	return submarine.multiply(), nil
}

func validCommand(s string) bool {
	switch s {
	case "forward", "down", "up":
		return true
	default:
		return false
	}
}

type coords struct {
	x, y, z int
}

func (c *coords) move(command string, distance int) {
	switch command {
	case "forward":
		c.x += distance
	case "down":
		c.y += distance
	case "up":
		c.y -= distance
	default:
		panic(fmt.Sprintf("bad input: %s", command))
	}
}

func (c *coords) moveWithAim(command string, distance int) {
	switch command {
	case "forward":
		c.x += distance
		c.y += c.z * distance
	case "down":
		c.z += distance
	case "up":
		c.z -= distance
	default:
		panic(fmt.Sprintf("bad input: %s", command))
	}
}

func (c *coords) multiply() int {
	return c.x * c.y
}
