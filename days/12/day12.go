package main

import (
	"aoc2021/utils/files"
	"fmt"
	"strings"
)

type caveNodes map[string][]string

func parseCaveNodes(input string) caveNodes {
	nodes := make(map[string][]string)

	for _, l := range strings.Split(input, "\n") {
		splitted := strings.Split(l, "-")
		from, to := splitted[0], splitted[1]
		nodes[from] = append(nodes[from], to)
		nodes[to] = append(nodes[to], from)
	}

	return nodes
}

func distinctPaths(nodes caveNodes, current string, visited map[string]bool) (count int) {
	if current == "end" {
		return 1
	}

	if _, ok := visited[current]; ok && current == strings.ToLower(current) {
		//small cave already visited
		return 0
	}

	visited[current] = true
	for _, to := range nodes[current] {

		clonedVisited := map[string]bool{}
		for k, v := range visited {
			clonedVisited[k] = v
		}

		count += distinctPaths(nodes, to, clonedVisited)
	}

	return count
}

func distinctPathsVisitTwice(nodes caveNodes, current, canVisitTwice, path string, alreadyVisitedTwice bool, visited map[string]bool) {
	currentPath := fmt.Sprintf("%s,%s", path, current)
	if current == "end" {
		visited[currentPath] = true
		return
	}

	if ok := strings.Contains(path, current); ok {
		if strings.ToLower(current) == current {
			if current != canVisitTwice || alreadyVisitedTwice {
				return
			} else {
				alreadyVisitedTwice = true
			}
		}
	}

	//visited[currentPath] = true
	for _, to := range nodes[current] {

		/*clonedVisited := map[string]bool{}
		for k, v := range visited {
			clonedVisited[k] = v
		}*/

		distinctPathsVisitTwice(
			nodes,
			to,
			canVisitTwice,
			currentPath,
			alreadyVisitedTwice,
			visited,
		)
	}

}

func getAllDistinctPaths(nodes caveNodes) int {
	var count int
	for _, edge := range nodes["start"] {
		count += distinctPaths(nodes, edge, map[string]bool{"start": true})
	}

	return count
}

func getAllDistinctVisitTwice(nodes caveNodes) int {
	solutions := map[string]bool{}

	for n := range nodes {
		if n == "start" || n == "end" || strings.ToLower(n) != n {
			continue
		}

		for _, edge := range nodes["start"] {
			distinctPathsVisitTwice(
				nodes,
				edge,
				n,
				"start",
				false,
				solutions,
			)
		}
	}

	return len(solutions)
}

func main() {
	puzzleInput := files.ReadInput()
	nodes := parseCaveNodes(puzzleInput)

	fmt.Println(getAllDistinctPaths(nodes))

	fmt.Println(getAllDistinctVisitTwice(nodes))

}
