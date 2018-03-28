package main

import (
	"DSA/ShortestPathAlgos/graph"
	"fmt"
	"log"
	"math"
)

func bellmanForAlgo(graphData graph.Graph, src int) {
	log.Println("Belmmand ford  algo starting", len(graphData.Edges))
	v := graphData.V
	dist := make([]int, v)

	for i := range dist {
		dist[i] = math.MaxInt64
	}

	dist[src] = 0

	for i := 1; i < v; i++ {
		for _, edge := range graphData.Edges {
			log.Println("@iteration withedge", edge.Src)
			if dist[edge.Src] != math.MaxInt64 && dist[edge.Src]+edge.Weight < dist[edge.Dest] {
				dist[edge.Dest] = dist[edge.Src] + edge.Weight
			}
		}
	}

	for _, edge := range graphData.Edges {
		if dist[edge.Src] != math.MaxInt64 && dist[edge.Src]+edge.Weight < dist[edge.Dest] {
			log.Println("Negative cycles")
		}
	}

	fmt.Printf("@heree %v", dist)
}

func main() {
	graphData := graph.CreateGraph(5, 8)
	graphData.Edges[0].Src = 0
	graphData.Edges[0].Dest = 1
	graphData.Edges[0].Weight = -1

	graphData.Edges[1].Src = 0
	graphData.Edges[1].Dest = 2
	graphData.Edges[1].Weight = 4
	graphData.Edges[2].Src = 1
	graphData.Edges[2].Dest = 2
	graphData.Edges[2].Weight = 3
	graphData.Edges[3].Src = 1
	graphData.Edges[3].Dest = 3
	graphData.Edges[3].Weight = 2
	graphData.Edges[4].Src = 1
	graphData.Edges[4].Dest = 4
	graphData.Edges[4].Weight = 2
	graphData.Edges[5].Src = 3
	graphData.Edges[5].Dest = 2
	graphData.Edges[5].Weight = 5
	graphData.Edges[6].Src = 3
	graphData.Edges[6].Dest = 1
	graphData.Edges[6].Weight = 1
	graphData.Edges[7].Src = 4
	graphData.Edges[7].Dest = 3
	graphData.Edges[7].Weight = -3
	bellmanForAlgo(graphData, 0)
}
