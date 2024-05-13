package main

import (
	"fmt"
	"os"
	"slices"
)

func printRoute(bestRoute [][]string, bestRouteInfo []int, trainNum int) {
	distributedTrains := distributeTrains(bestRouteInfo, trainNum)
	printTrainTurns(distributedTrains, bestRouteInfo, bestRoute, trainNum)
}

func distributeTrains(routeInfo []int, trainNum int) map[int][]int {
	turn := 1
	distributedTrainNum := 0
	distributedTrains := make(map[int][]int, len(routeInfo))

	for distributedTrainNum < trainNum {
		for index, routeNum := range routeInfo {
			if routeNum <= turn {
				distributedTrainNum++
				if distributedTrainNum <= trainNum {
					trainName := distributedTrainNum
					distributedTrains[index] = append(distributedTrains[index], trainName)
				}
			}
		}
		turn++
	}
	return distributedTrains
}

func printTrainTurns(distributedTrains map[int][]int, bestRouteInfo []int, bestRoutes [][]string, trainNum int) {
	// Make a map to keep track of the stations' statuses
	stationStatus := make(map[string]int) // options: 0 = available, 1 = occupied
	for _, stations := range bestRoutes {
		for _, station := range stations {
			stationStatus[station] = 0
		}
	}

	// Make a map to keep track of the trains' statuses
	type trainStatus struct {
		status               string // options: "starting", "moving", "finished"
		pathNumber           int
		currentStationNumber int
	}

	trainsStatusMap := make(map[int]*trainStatus)

	// Function for finding which path the train is supposed to go
	pathNumberFunc := func(i int) int {
		var path int
		for path = 0; path < len(bestRouteInfo); path++ {
			if slices.Contains(distributedTrains[path], i) {
				break
			}
		}
		return path
	}

	for train := 1; train <= trainNum; train++ {
		trainData := &trainStatus{
			status:               "starting",
			pathNumber:           pathNumberFunc(train),
			currentStationNumber: 0,
		}
		trainsStatusMap[train] = trainData
	}

	// Main loop to simulate trains' movements, update stations' and trains' statuses and print the train turns
	var trainTurn string
	var oneLengthPathUsed bool
	endStation := os.Args[3]
	for trainsStatusMap[trainNum].status != "finished" {
		for Ti := 1; Ti <= trainNum; Ti++ {
			if trainsStatus, ok := trainsStatusMap[Ti]; ok {
				if trainsStatus.status == "moving" {
					nextStation := bestRoutes[trainsStatus.pathNumber][trainsStatus.currentStationNumber+1]
					if stationStatus[nextStation] == 0 {
						currentStation := bestRoutes[trainsStatus.pathNumber][trainsStatus.currentStationNumber]
						nextStation := bestRoutes[trainsStatus.pathNumber][trainsStatus.currentStationNumber+1]
						stationStatus[currentStation] = 0
						trainsStatus.currentStationNumber++
						trainTurn = trainTurn + fmt.Sprintf("T%d-%v ", Ti, nextStation)
						if nextStation == endStation {
							trainsStatus.status = "finished"
							stationStatus[nextStation] = 0
						} else {
							stationStatus[nextStation] = 1
						}
					}

				} else if trainsStatus.status == "starting" {
					startingStation := bestRoutes[trainsStatus.pathNumber][0]
					if stationStatus[startingStation] == 0 {
						if startingStation == endStation {
							if !oneLengthPathUsed {
								trainsStatus.status = "finished"
								oneLengthPathUsed = true
								trainTurn = trainTurn + fmt.Sprintf("T%d-%v ", Ti, startingStation)
							}
							continue
						} else {
							stationStatus[startingStation] = 1
							trainsStatus.status = "moving"
						}
						trainTurn = trainTurn + fmt.Sprintf("T%d-%v ", Ti, startingStation)
					}
				}
			}
		}

		fmt.Println(trainTurn)
		trainTurn = ""
		oneLengthPathUsed = false
	}
}
