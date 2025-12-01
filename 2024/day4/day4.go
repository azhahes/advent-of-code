package day4

import (
	"bufio"
	"fmt"

	"github.com/azhahes/advent-of-code/util/file"
)

const (
	up          = "up"
	down        = "down"
	right       = "right"
	left        = "left"
	topLeft     = "topLeft"
	topRight    = "topRight"
	bottomLeft  = "bottomLeft"
	bottomRight = "bottomRight"
)

var xmas = []byte{'X', 'M', 'A', 'S'}

func boundryCheck(i, j int, matrix []string) bool {
	if i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[0]) {
		return true
	}
	return false
}

func traverse(i, j int, matrix []string) (byte, error) {
	if boundryCheck(i, j, matrix) {
		return matrix[i][j], nil
	}
	return 0, fmt.Errorf("Out of Range")
}

var directionsfn = map[string]func(int, int, int, []string) (byte, error){
	up:          func(i, j, k int, matrix []string) (byte, error) { return traverse(i, j-k, matrix) },
	down:        func(i, j, k int, matrix []string) (byte, error) { return traverse(i, j+k, matrix) },
	right:       func(i, j, k int, matrix []string) (byte, error) { return traverse(i+k, j, matrix) },
	left:        func(i, j, k int, matrix []string) (byte, error) { return traverse(i-k, j, matrix) },
	topLeft:     func(i, j, k int, matrix []string) (byte, error) { return traverse(i-k, j-k, matrix) },
	topRight:    func(i, j, k int, matrix []string) (byte, error) { return traverse(i+k, j-k, matrix) },
	bottomLeft:  func(i, j, k int, matrix []string) (byte, error) { return traverse(i-k, j+k, matrix) },
	bottomRight: func(i, j, k int, matrix []string) (byte, error) { return traverse(i+k, j+k, matrix) },
}

func dfs(i, j int, matrix []string) int {
	result := 0
	for _, nextLetter := range directionsfn {
		k := 1
		for l := 1; l < len(xmas); l++ {
			letter, err := nextLetter(i, j, l, matrix)
			if err != nil {
				break
			}
			if letter == xmas[l] {
				k++
			}
		}
		if k == len(xmas) {
			result++
		}
	}
	return result
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
