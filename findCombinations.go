package main

import (
	"os"
	"sort"
)

func findRouteCombinations(allRoutes [][]string) [][][]string {
	routesNum := len(allRoutes)
	routes := make([][][]string, 0)

	for startNum := range routesNum {
		var combination [][]string
		combination = append(combination, allRoutes[startNum])
		addNewcombination(allRoutes, combination, routesNum, startNum+1, &routes)
	}
	return routes
}

func addNewcombination(allRoutes [][]string, combination [][]string, endNum int, countNum int, routes *[][][]string) {
	if countNum == endNum {
		if checkRedundancy(combination, routes) && combination != nil {
			*routes = append(*routes, combination)
		}
	} else {
		if checkAdditionPossibility(allRoutes[countNum], combination) {

			newCombination := append(combination, allRoutes[countNum])
			addNewcombination(allRoutes, newCombination, endNum, countNum+1, routes)

		}
		addNewcombination(allRoutes, combination, endNum, countNum+1, routes)
	}
}

func checkRedundancy(combination [][]string, routes *[][][]string) bool {
	totalLength := len(combination)
	var routesLengthList [][]int
	var combinationLengthList []int

	if totalLength == 1 {
		routeLength := len(combination[0])
		for _, routeSet1 := range *routes {
			for _, route1 := range routeSet1 {
				if len(route1) == routeLength {
					return false
				}
			}
		}
	} else {
		for _, routeSet2 := range *routes {
			var tempLengthList []int
			for _, route1 := range routeSet2 {
				tempLengthList = append(tempLengthList, len(route1))
			}
			routesLengthList = append(routesLengthList, tempLengthList)
		}

		for _, route2 := range combination {
			combinationLengthList = append(combinationLengthList, len(route2))
		}

		for _, routeLengthList := range routesLengthList {
			if checkInclusiveCombination(combinationLengthList, routeLengthList) {
				return false
			}
		}
	}
	return true
}

func checkInclusiveCombination(combination, route []int) bool {
	if len(combination) > len(route) {
		return false
	}
	sort.Ints(combination)
	sort.Ints(route)

	for index, number := range combination {
		if number != route[index] {
			return true
		}
	}
	return false
}

func checkAdditionPossibility(checkList []string, combination [][]string) bool {
	endStation := os.Args[3]
	for _, stationlist := range combination {
		for _, station1 := range stationlist {
			for _, station2 := range checkList {
				if station1 == endStation || station2 == endStation {
					continue
				} else if station1 == station2 {
					return false
				}
			}
		}
	}
	return true
}
