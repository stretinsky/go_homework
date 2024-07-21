package main

import (
	"fmt"
)

func EvalSequence(matrix [][]int, userAnswer []int) int {
	// validation
	if err := validateInput(matrix, userAnswer); err != nil {
		return -1
	}

	maxGrade := calMaxGrade(matrix)
	userGrade := calcUserGrade(matrix, userAnswer)

	if maxGrade == 0 {
		return 0
	}
	percent := userGrade * 100 / maxGrade

	return percent
}

func calMaxGrade(matrix [][]int) int {
	n := len(matrix)
	maxDistance := 0

	var dfs func(node int, visited []bool, currentDistance int) int
	dfs = func(node int, visited []bool, currentDistance int) int {
		visited[node] = true
		maxDist := currentDistance

		for i := 0; i < n; i++ {
			if matrix[node][i] != 0 && !visited[i] {
				dist := dfs(i, visited, currentDistance+matrix[node][i])
				if dist > maxDist {
					maxDist = dist
				}
			}
		}

		visited[node] = false
		return maxDist
	}

	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		distance := dfs(i, visited, 0)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func calcUserGrade(matrix [][]int, userAnswer []int) int {
	userGrade := 0
	for i := 1; i < len(userAnswer); i++ {
		userGrade += matrix[userAnswer[i-1]][userAnswer[i]]
	}
	return userGrade
}

func validateInput(matrix [][]int, userAnswer []int) error {
	n := len(matrix)

	for _, row := range matrix {
		if len(row) != n {
			return fmt.Errorf("matrix must be square")
		}
	}

	for i := 0; i < n; i++ {
		if matrix[i][i] != 0 {
			return fmt.Errorf("there should be no loops in the matrix")
		}
	}

	seen := make(map[int]bool)
	for _, answer := range userAnswer {
		if answer < 0 || answer >= n {
			return fmt.Errorf("user's responses should not exceed the range of the matrix")
		}
		if seen[answer] {
			return fmt.Errorf("elements in the user's response slice must be unique")
		}
		seen[answer] = true
	}
	return nil
}

