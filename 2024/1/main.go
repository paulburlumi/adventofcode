package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	if err := run(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func run() error {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var left, right []int
	for scanner.Scan() {
		line := scanner.Text()
		var a, b int
		_, err := fmt.Sscanf(line, "%d %d", &a, &b)
		if err != nil {
			return fmt.Errorf("parsing line %q: %w", line, err)
		}
		left = append(left, a)
		right = append(right, b)
	}
	slices.Sort(left)
	slices.Sort(right)
	if len(left) != len(right) {
		return fmt.Errorf("slices have different lengths: %d != %d", len(left), len(right))
	}
	distance := 0
	for i := range left {
		if left[i] > right[i] {
			distance += left[i] - right[i]
		} else {
			distance += right[i] - left[i]
		}
	}
	fmt.Printf("distance: %d\n", distance)

	counts := map[int]int{}
	for i := range right {
		counts[right[i]]++
	}
	similarity := 0
	for i := range left {
		similarity += left[i] * counts[left[i]]
	}
	fmt.Printf("similarity: %d\n", similarity)

	return nil
}
