package day1

import (
	"bufio"
	"math"
	"sort"
	"strings"

	"github.com/azhahes/advent-of-code/2024/file"
	"github.com/azhahes/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	f := file.Open(filePath)
	left, right := make([]int, 0), make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		left = append(left, parse.Int(words[0]))
		right = append(right, parse.Int(words[1]))
	}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i, v := range left {
		diff := int(math.Abs(float64(v - right[i])))
		sum += diff
	}
	return sum
}

func part2(filePath string) int {
	f := file.Open(filePath)
	left, rightCount := make([]int, 0), make(map[int]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		left = append(left, parse.Int(words[0]))
		right := parse.Int(words[1])
		if _, ok := rightCount[right]; ok {
			rightCount[right]++
		} else {
			rightCount[right] = 1
		}
	}
	sum := 0
	for _, v := range left {
		if val, ok := rightCount[v]; ok {
			sum += val * v
		}
	}
	return sum
}
