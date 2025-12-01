package day3

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/azhahes/advent-of-code/util/file"
	"github.com/azhahes/advent-of-code/util/parse"
)

func part1(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		for _, match := range re.FindAllSubmatch(scanner.Bytes(), -1) {
			a, b := parse.MustToInt(string(match[1])), parse.MustToInt(string(match[2]))
			sum += (a * b)
		}
	}
	return sum
}

func part2(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	sum := 0
	do := true
	for scanner.Scan() {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)
		for _, match := range re.FindAllSubmatch(scanner.Bytes(), -1) {
			if strings.Contains(string(match[0]), "mul") && do {
				a, b := parse.MustToInt(string(match[1])), parse.MustToInt(string(match[2]))
				sum += (a * b)
			} else if string(match[0]) == "don't()" {
				do = false
			} else if string(match[0]) == "do()" {
				do = true
			}
		}
	}
	return sum
}
