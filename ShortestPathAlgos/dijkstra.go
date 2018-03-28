package main

import (
	"fmt"
	"log"
)

func minDistance(dist [9]int, sptSet [9]bool) int {
	const MaxInt64 = 1<<63 - 1
	var min, minIndex int
	min = MaxInt64
	for v := 0; v < 9; v++ {
		if !sptSet[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}
	return minIndex
}

func printArray(dist [9]int) {
	fmt.Printf("@heree %v", dist)
}
func dijikstra(graph [9][9]int, src int) {
	var dist [9]int
	var sptSet [9]bool
	const MaxInt64 = 1<<63 - 1
	for i := 0; i < 9; i++ {
		dist[i] = MaxInt64
		sptSet[i] = false
	}
	dist[src] = 0
	for count := 0; count < 8; count++ {
		u := minDistance(dist, sptSet)
		sptSet[u] = true
		for v := 0; v < 9; v++ {
			if (!sptSet[v]) && graph[u][v] != 0 &&
				(dist[u] != MaxInt64) && (dist[u]+graph[u][v] < dist[v]) {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}
	printArray(dist)
}
func main() {
	log.Println("Here")
	graph := [9][9]int{{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}
	dijikstra(graph, 0)
}
