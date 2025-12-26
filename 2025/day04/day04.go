package day04

import (
	"bufio"

	"github.com/azhahes/advent-of-code/util/file"
)

func Part1(filePath string) int {
	file := file.MustOpen(filePath)
	scanner := bufio.NewScanner(file)
	wall := make([][]rune, 0)
	result := 0
	for scanner.Scan() {
		row := scanner.Text()
		objects := make([]rune, len(row))
		for _, r := range row {
			objects = append(objects, r)
		}
		wall = append(wall, objects)
	}
	for i, row := range wall {
		for j := range row {
			if wall[i][j] == '@' && countAdj(i, j, wall) < 4 {
				result++
			}
		}
	}
	return result
}

func countAdj(i, j int, wall [][]rune) int {
	r, c := len(wall), len(wall[0])
	count := 0
	directions := [][]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, 1},   // right
		{0, -1},  // left
		{1, 1},   // bottom right
		{1, -1},  // bottom left
		{-1, -1}, // top left
		{-1, 1},  // top bottom
	}

	for _, dir := range directions {
		newi, newj := i+dir[0], j+dir[1]
		if newi < r && newi >= 0 && newj < c && newj >= 0 && wall[newi][newj] == '@' {
			count++
		}
	}
	return count
}

func Part2(filePath string) int {
	file := file.MustOpen(filePath)
	scanner := bufio.NewScanner(file)
	wall := make([][]rune, 0)
	result := 0
	for scanner.Scan() {
		row := scanner.Text()
		objects := make([]rune, len(row))
		for _, r := range row {
			objects = append(objects, r)
		}
		wall = append(wall, objects)
	}
	for {
		curr := 0
		visited := make([][]int, 0)
		for i, row := range wall {
			for j := range row {
				if wall[i][j] == '@' && countAdj(i, j, wall) < 4 {
					curr++
					visited = append(visited, []int{i, j})
				}
			}
		}
		for _, v := range visited {
			wall[v[0]][v[1]] = '.'
		}
		if curr == 0 {
			break
		}
		result += curr
	}
	return result
}
