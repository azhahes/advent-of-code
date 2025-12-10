package day1

import (
	"bufio"

	"github.com/azhahes/advent-of-code/util/file"
	"github.com/azhahes/advent-of-code/util/parse"
)

func part1(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	dialPos := 50
	password := 0
	for scanner.Scan() {
		line := scanner.Text()
		side := rune(line[0])
		rotation := parse.MustToInt(line[1:])
		if side == 'L' {
			dialPos -= (rotation % 100)
		} else {
			dialPos += (rotation % 100)
		}

		if dialPos < 0 {
			dialPos += 100
		} else if dialPos > 99 {
			dialPos -= 100
		}

		if dialPos == 0 {
			password++
		}

	}
	return password
}

func part2(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	dialPos := 50
	password := 0
	for scanner.Scan() {
		line := scanner.Text()
		side := rune(line[0])
		rotation := parse.MustToInt(line[1:])
		password += rotation / 100
		if side == 'L' {
			dialPos -= (rotation % 100)
		} else {
			dialPos += (rotation % 100)
		}
		if dialPos < 0 {
			dialPos += 100
			password++
		} else if dialPos > 99 {
			dialPos -= 100
			password++
		}

	}
	return password
}
