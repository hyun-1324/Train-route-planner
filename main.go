package main

import (
	"fmt"
	"os"
)

func main() {
	trainNum, stations, connections, err := validateInputAndPickDatafromNetWorkMap()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	connectionsMap := inputDataToMap(stations, connections)

	start := os.Args[2]
	end := os.Args[3]

	allRoutes, err := findAllRoutes(connectionsMap, start, end)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	combinationRoutes := findRouteCombinations(allRoutes)

	bestRoute, bestRouteInfo := findBestRouteByTrainNumber(trainNum, combinationRoutes)

	printRoute(bestRoute, bestRouteInfo, trainNum)
}
