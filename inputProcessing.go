package main

// Function to process valid input data to the format used in the algorithm
func inputDataToMap(stations []station, connections [][]string) map[string][]string {
	connectionsMap := make(map[string][]string)
	var connectedStations []string

	for _, station := range stations {
		for _, connectionPair := range connections {
			if connectionPair[0] == station.name {
				connectedStations = append(connectedStations, connectionPair[1])
			} else if connectionPair[1] == station.name {
				connectedStations = append(connectedStations, connectionPair[0])
			}
		}
		if connectedStations != nil {
			connectionsMap[station.name] = connectedStations
			connectedStations = nil
		}
	}

	return connectionsMap
}
