package main

import "math"

func findBestRouteByTrainNumber(trainNumber int, combinationRoutes [][][]string) (bestRoute [][]string, bestRouteInfo []int) {
	CombinationRoutesInfo := getCombinationRoutesInfo(combinationRoutes)
	shortestTurn := math.MaxInt

	for index, combinationRouteInfo := range CombinationRoutesInfo {
		turnRecord := findShortestTurn(combinationRouteInfo, trainNumber)
		if shortestTurn > turnRecord {
			bestRoute = combinationRoutes[index]
			bestRouteInfo = combinationRouteInfo
			shortestTurn = turnRecord
		}
	}
	return bestRoute, bestRouteInfo
}

func getCombinationRoutesInfo(combinationRoutes [][][]string) [][]int {
	var routesInfo [][]int

	for _, combinationRoute := range combinationRoutes {
		var routeInfo []int
		for _, route := range combinationRoute {
			routeInfo = append(routeInfo, len(route))
		}
		routesInfo = append(routesInfo, routeInfo)
	}
	return routesInfo
}

func findShortestTurn(routeInfo []int, trainNumber int) (shortestTurn int) {
	trainCounter := 0
	turn := 0

	for trainCounter < trainNumber {
		trainCounter, turn = calculateTrainsSent(trainCounter, turn, routeInfo)
	}
	return turn
}

func calculateTrainsSent(trainCounter int, turn int, routeInfo []int) (int, int) {
	plusNum := 0
	for _, routeNum := range routeInfo {
		if routeNum <= turn {
			plusNum++
		}
	}

	trainCounter = trainCounter + plusNum

	return trainCounter, turn + 1

}
