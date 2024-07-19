package main

import (
	"fmt"
)


const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func bfs(matrix [][]int, start int, goal int) int {
	n := len(matrix)

	// посещенные узлы
	var visited = make(map[int]bool, n)
	// расстояния до узлов
	var distances = make([]int, n)
	for i := 0; i < n; i++ {
		distances[i] = MaxInt
	}
	distances[start] = 0

	queue := make(chan int, n*n)
	queue <- start

	for len(queue) > 0 {
		node := <-queue
		if node == goal {
			return distances[node]
		}
		for i := 0; i < n; i++ {
			if matrix[node][i] != 0 && !visited[i] {
				queue <- i
				newDistance := distances[node] + matrix[node][i]
				if newDistance < distances[i] {
					distances[i] = newDistance
					queue <- i
				}
			}
		}
		visited[node] = true
	}
	return distances[goal]
}


func main() {
	// visual representation in the readme
	matrix := [][]int{
		{0, 1, 4, 0},
		{1, 0, 2, 6},
		{4, 2, 0, 3},
		{0, 6, 3, 0},
	}

	fmt.Println(bfs(matrix, 0, 3))
}
