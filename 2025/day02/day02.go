package day02

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/azhahes/advent-of-code/util/file"
	"github.com/azhahes/advent-of-code/util/parse"
)

func Part1(filePath string) int64 {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	var sum int64
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}
	ranges := strings.Split(line, ",")
	for _, r := range ranges {
		numbers := strings.Split(r, "-")
		start := parse.MustToInt(numbers[0])
		end := parse.MustToInt(numbers[1])
		for i := start; i <= end; i++ {
			numString := strconv.Itoa(i)
			n := len(numString)
			if n%2 == 0 {
				j, k := 0, n/2
				for k < n {
					if numString[j] != numString[k] {
						break
					}
					j++
					k++
				}
				if k == n {
					sum += int64(parse.MustToInt(numString))
				}
			}
		}
	}
	return sum
}

func Part2(filePath string) int64 {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	var sum int64
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}
	ranges := strings.Split(line, ",")
	for _, r := range ranges {
		numbers := strings.Split(r, "-")
		start := parse.MustToInt(numbers[0])
		end := parse.MustToInt(numbers[1])
		for i := start; i <= end; i++ {
			numString := strconv.Itoa(i)
			n := len(numString) / 2
			for j := range n {
				firstPart := strings.Index(numString[1:], string(numString[:j]))
				if firstPart != -1 && strings.Join(strings.Split(numString, numString[:firstPart+1]), "") == "" {
					// fmt.Printf("%s : %d\n", numString, firstPart)
					sum += int64(i)
					break
				}
			}

		}
	}
	return sum
}
