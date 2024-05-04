package main

import (
	"fmt"
	"slices"
	"sort"
)

func findAllRoutes(connections map[string][]string, start string, end string) ([][]string, error) {

	var allPaths [][]string
	var path []string
	findAllRoutesByRecursion(connections, start, end, path, &allPaths)
	if allPaths == nil {
		err := fmt.Errorf("no path exists between %s and %s", start, end)
		return nil, err
	}

	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	return allPaths, nil
}

func findAllRoutesByRecursion(connections map[string][]string, start string, end string, path []string, allPaths *[][]string) {
	path = append(path, start)

	if start == end {
		tempPath := make([]string, len(path))
		copy(tempPath, path)
		*allPaths = append(*allPaths, tempPath[1:])
	} else {
		for _, neighbor := range connections[start] {
			if !slices.Contains(path, neighbor) {
				findAllRoutesByRecursion(connections, neighbor, end, path, allPaths)
			}
		}
	}
}
