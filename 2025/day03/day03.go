package day03

import (
	"bufio"

	"github.com/azhahes/advent-of-code/util/file"
	"github.com/azhahes/advent-of-code/util/parse"
)

func Part1(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		first, second := line[0], line[1]
		i := 0
		for i < len(line)-1 {
			if line[i] > first {
				first = line[i]
				second = line[i+1]
			} else if line[i] > second {
				second = line[i]
			}
			i++
		}
		if second < line[len(line)-1] {
			second = line[len(line)-1]
		}
		sum += parse.MustToInt(string(first) + string(second))
	}
	return sum
}

func Part2(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		prevStart := 0
		end := len(line) - 12
		value := 0
		/*
		   234234234234278
		*/
		for end < len(line) {
			val, index := findGreatestDigit(prevStart, end, line)
			value *= 10
			value += val
			prevStart = index + 1
			end++
		}
		sum += value
	}
	return sum
}

func findGreatestDigit(start, end int, line string) (int, int) {
	v, i := 0, start
	for start <= end {
		parsedVal := parse.MustToInt(string(line[start]))
		if v < parsedVal {
			v = parsedVal
			i = start
		}
		start++
	}
	return v, i
}
