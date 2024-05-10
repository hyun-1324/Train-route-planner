package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func validateInputAndPickDatafromNetWorkMap() (int, []station, [][]string, error) {
	err := checkArguments()
	if err != nil {
		return 0, nil, nil, err
	}

	numTrains, err := checkNumberOfTrains()
	if err != nil {
		return 0, nil, nil, err
	}

	netWorkMapLines, err := netWorkMapToSliceOfStrings()
	if err != nil {
		return 0, nil, nil, err
	}

	stationsList, connectionsList, err := extractStationsAndConnections(netWorkMapLines)
	if err != nil {
		return 0, nil, nil, err
	}

	validatedStations, err := validateStationData(stationsList)
	if err != nil {
		return 0, nil, nil, err
	}

	if len(validatedStations) > 10000 {
		err := fmt.Errorf("too many stations in the network map (over 10000): %v", len(validatedStations))
		return 0, nil, nil, err
	}

	err = checkIfStationsExist(validatedStations)
	if err != nil {
		return 0, nil, nil, err
	}

	validatedConnections, err := validateConnectionData(connectionsList, validatedStations)
	if err != nil {
		return 0, nil, nil, err
	}

	return numTrains, validatedStations, validatedConnections, nil
}

func checkArguments() error {
	if numArgs := len(os.Args); numArgs != 5 {
		err := fmt.Errorf("incorrect number of command line arguments: %v", numArgs)
		return err
	}

	if os.Args[2] == os.Args[3] { // start station == end station
		err := fmt.Errorf("the start station is the same as end station: %v", os.Args[2])
		return err
	}

	return nil
}

func checkNumberOfTrains() (int, error) {
	numTrainsStr := os.Args[4]
	numTrainsInt, err := strconv.Atoi(numTrainsStr)
	if err != nil {
		err = fmt.Errorf("number of trains in not a valid number: %v", numTrainsStr)
		return 0, err
	}

	if numTrainsInt < 1 {
		err := fmt.Errorf("number of trains is negative or 0: %v", numTrainsStr)
		return 0, err
	}

	return numTrainsInt, nil
}

func netWorkMapToSliceOfStrings() ([]string, error) {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		err = fmt.Errorf("network map not found from path: %v", filePath)
		return nil, err
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var netWorkMapLines []string

	for fileScanner.Scan() {
		lineStr := fileScanner.Text()
		// Drop blank lines and lines including only comment
		if lineStr != "" && lineStr[0] != '#' {
			// Remove comments if they are on the same line as valid data
			dataAndComment := strings.Split(lineStr, "#")
			lineStr = dataAndComment[0]
			lineStr = strings.TrimSpace(lineStr)
			netWorkMapLines = append(netWorkMapLines, lineStr)
		}
	}

	if err := fileScanner.Err(); err != nil {
		err = fmt.Errorf("error scanning the network map file: %v", filePath)
		return nil, err
	}

	return netWorkMapLines, nil
}

func extractStationsAndConnections(netWorkMapLines []string) ([]string, []string, error) {
	var stations, connections []string

	startStations := slices.Index(netWorkMapLines, "stations:")
	startConnections := slices.Index(netWorkMapLines, "connections:")

	// Check for missing declarations
	if startStations == -1 {
		err := fmt.Errorf("'stations:' declaration is missing from the network map")
		return nil, nil, err
	}

	if startConnections == -1 {
		err := fmt.Errorf("'connections:' declaration is missing from the network map")
		return nil, nil, err
	}

	// Check for missing data after declarations
	if startStations == len(netWorkMapLines)-1 || netWorkMapLines[startStations+1] == "connections:" {
		err := fmt.Errorf("there are no stations listed in the network map")
		return nil, nil, err
	}

	if startConnections == len(netWorkMapLines)-1 || netWorkMapLines[startConnections+1] == "stations:" {
		err := fmt.Errorf("there are no connections listed in the network map")
		return nil, nil, err
	}

	// Extract stations
	for _, row := range netWorkMapLines[startStations+1:] {
		if row == "connections:" {
			break
		}
		stations = append(stations, row)
	}

	// Extract connections
	for _, row := range netWorkMapLines[startConnections+1:] {
		if row == "stations:" {
			break
		}
		connections = append(connections, row)
	}

	return stations, connections, nil
}

type station struct {
	name string
	x    int
	y    int
}

func validateStationData(stationsList []string) ([]station, error) {
	var validStations []station

	for _, row := range stationsList {
		stationData := strings.Split(row, ",")
		if len(stationData) != 3 {
			err := fmt.Errorf("incorrect data in the stations section of the network map: %v", row)
			return nil, err
		}

		stationData[0] = strings.TrimSpace(stationData[0])
		stationData[1] = strings.TrimSpace(stationData[1])
		stationData[2] = strings.TrimSpace(stationData[2])

		// Check station name validity
		stationName := stationData[0]
		for _, character := range stationName {

			if (character < 'a' || character > 'z') && (character < '0' || character > '9') && character != '_' {
				err := fmt.Errorf("invalid station name in the network map: %v", stationName)
				return nil, err
			}
		}

		// Check validity of coordinates
		xCoordinate, err := strconv.Atoi(stationData[1])
		if err != nil {
			err := fmt.Errorf("x-coordinate is not an integer: %v", row)
			return nil, err
		}

		if xCoordinate < 0 {
			err := fmt.Errorf("x-coordinate is negative: %v", row)
			return nil, err
		}

		yCoordinate, err := strconv.Atoi(stationData[2])
		if err != nil {
			err := fmt.Errorf("y-coordinate is not an integer: %v", row)
			return nil, err
		}

		if yCoordinate < 0 {
			err := fmt.Errorf("y-coordinate is negative: %v", row)
			return nil, err
		}

		// Append data to list of valid stations while checking for duplicates in the data
		validStationData := station{
			name: stationData[0],
			x:    xCoordinate,
			y:    yCoordinate,
		}

		for _, existingStation := range validStations {
			if existingStation.name == validStationData.name {
				err := fmt.Errorf("duplicate station name in the network map: %v", existingStation.name)
				return nil, err
			}
			if (existingStation.x == validStationData.x) && (existingStation.y == validStationData.y) {
				err := fmt.Errorf("stations have same coordinates: %v & %v", existingStation.name, validStationData.name)
				return nil, err
			}
		}
		validStations = append(validStations, validStationData)
	}

	return validStations, nil
}

func checkIfStationsExist(validatedSations []station) error {
	var startExists, endExists bool

	startStation := os.Args[2]
	endStation := os.Args[3]

	for _, stationData := range validatedSations {
		if stationData.name == startStation {
			startExists = true
		}
		if stationData.name == endStation {
			endExists = true
		}
	}

	if !startExists {
		err := fmt.Errorf("the start station %v does not exist in the network map", startStation)
		return err
	}
	if !endExists {
		err := fmt.Errorf("the end station %v does not exist in the network map", endStation)
		return err
	}

	return nil
}

func validateConnectionData(connectionsList []string, validatedSations []station) ([][]string, error) {
	var validConnectionPairs [][]string

	for _, pairStr := range connectionsList {
		connectionPair := strings.Split(pairStr, "-")
		if len(connectionPair) != 2 {
			err := fmt.Errorf("incorrect formatting of the connection data in the network map: %v", pairStr)
			return nil, err
		}

		connectionPair[0] = strings.TrimSpace(connectionPair[0])
		connectionPair[1] = strings.TrimSpace(connectionPair[1])

		// Validate station names in connection pairs
		var validStationName bool

		for _, stationName := range connectionPair {
			for _, stationData := range validatedSations {
				if stationData.name == stationName {
					validStationName = true
				}
			}
			if !validStationName {
				err := fmt.Errorf("station used in the connections does not exist: %v", stationName)
				return nil, err
			} else {
				validStationName = false
			}
		}

		slices.Sort(connectionPair)

		// Add slice to list of pairs while checking for duplicates
		for _, existingPair := range validConnectionPairs {
			if slices.Equal(connectionPair, existingPair) {
				err := fmt.Errorf("double connection between %v & %v in the network map", existingPair[0], existingPair[1])
				return nil, err
			}
		}
		validConnectionPairs = append(validConnectionPairs, connectionPair)
	}

	return validConnectionPairs, nil
}
