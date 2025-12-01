package day5

import (
	"bufio"
	"strings"

	"github.com/azhahes/advent-of-code/2024/file"
	"github.com/azhahes/advent-of-code/2024/parse"
)

type pages struct {
	nextPages []int
	prevPages []int
}

var (
	rules        map[int]*pages
	orderedPages []pages
)

func parseRulesAndPages(filepath string) {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			break
		}
		rule := strings.Split(scanner.Text(), "|")
		prevPage, nextPage := parse.MustToInt(rule[0]), parse.MustToInt(rule[1])
		rules[prevPage].nextPages = append(rules[prevPage].nextPages, nextPage)
		rules[nextPage].prevPages = append(rules[nextPage].prevPages, prevPage)
	}

	for scanner.Scan() {
		orderedPages := parse.MustToIntSlice(strings.Split(scanner.Text(), ","))
	}
}

func part1(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	result := 0
	matrix := make([]string, 0)
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'X' {
				result += dfs(i, j, matrix)
			}
		}
	}

	return result
}

func part2(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	result := 0
	matrix := make([]string, 0)
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'A' {
				trVal, err := directionsfn[topRight](i, j, 1, matrix)
				if err != nil {
					continue
				}
				tlVal, err := directionsfn[topLeft](i, j, 1, matrix)
				if err != nil {
					continue
				}
				brVal, err := directionsfn[bottomRight](i, j, 1, matrix)
				if err != nil {
					continue
				}
				blVal, err := directionsfn[bottomLeft](i, j, 1, matrix)
				if err != nil {
					continue
				}
				if ((trVal == 'M' && blVal == 'S') || (trVal == 'S' && blVal == 'M')) && ((tlVal == 'M' && brVal == 'S') || (tlVal == 'S' && brVal == 'M')) {
					result++
				}
			}
		}
	}
	return result
}
