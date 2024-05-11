package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	type inputOutputData struct {
		arguments    []string
		maxLenOutput int
		mapName      string
	}

	inputOutputList := []inputOutputData{{
		arguments:    []string{"./test_suite/londonNetworkMap.txt", "waterloo", "st_pancras", "2"},
		maxLenOutput: 2,
		mapName:      "London Network Map",
	}, {
		arguments:    []string{"./test_suite/fairylandNetworkMap.txt", "jungle", "desert", "10"},
		maxLenOutput: 8,
		mapName:      "Fairyland Network Map",
	}, {
		arguments:    []string{"./test_suite/foodNetworkMap.txt", "bond_square", "space_port", "4"},
		maxLenOutput: 6,
		mapName:      "Food Network Map",
	}, {
		arguments:    []string{"./test_suite/composersNetworkMap.txt", "beethoven", "part", "9"},
		maxLenOutput: 6,
		mapName:      "Composers network Map",
	}, {
		arguments:    []string{"./test_suite/distanceNetworkMap.txt", "beginning", "terminus", "20"},
		maxLenOutput: 11,
		mapName:      "Distance Network Map",
	}, {
		arguments:    []string{"./test_suite/numbersNetworkMap.txt", "two", "four", "4"},
		maxLenOutput: 6,
		mapName:      "Numbers Network Map",
	}, {
		arguments:    []string{"./test_suite/sizeNetworkMap.txt", "small", "large", "9"},
		maxLenOutput: 8,
		mapName:      "Size Network Map",
	},
	}

	reset := "\x1B[0m"
	red := "\x1B[1;31m"
	green := "\x1B[1;32m"
	blue := "\x1B[1;34m"
	yellow := "\x1B[1;33m"
	bold := "\x1B[1m"

	var summary []string
	var result string

	fmt.Printf("____________________\n\n")

	for _, testCase := range inputOutputList {
		cmd := exec.Command("./trains", testCase.arguments...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Errorf("%sError running program:%s %s", red, reset, err)
			continue
		}
		if string(output)[:6] == "Error:" {
			result = fmt.Sprintf("%sINVALID INPUT%s %s%s: %s trains between %s and %s%s", yellow, reset, bold, testCase.mapName, testCase.arguments[3], testCase.arguments[1], testCase.arguments[2], reset)
			fmt.Printf("%s\nProgram reported an error:\n%s____________________\n\n", result, string(output))
			summary = append(summary, result)
			continue
		}

		if testOtputLen := len(strings.Split(string(output), "\n")) - 1; testOtputLen <= testCase.maxLenOutput {
			result = fmt.Sprintf("%sPASSED%s %s%s: %s trains between %s and %s%s", green, reset, bold, testCase.mapName, testCase.arguments[3], testCase.arguments[1], testCase.arguments[2], reset)
			fmt.Println(result)
			fmt.Printf("Number of train turns used: %d\nMaximum train turns usable: %d\n\n", testOtputLen, testCase.maxLenOutput)
			fmt.Printf("Train turns:\n%s____________________\n\n", string(output))
		} else {
			result = fmt.Sprintf("%sFAILED%s %s%s: %s trains between %s and %s%s", red, reset, bold, testCase.mapName, testCase.arguments[3], testCase.arguments[1], testCase.arguments[2], reset)
			trainTurns := fmt.Sprintf("Number of train turns used: %d\nMaximum train turns usable: %d\n\n", testOtputLen, testCase.maxLenOutput)
			terminalOutput := fmt.Sprintf("Train turns:\n%s____________________\n\n", string(output))
			fmt.Printf("%s\n%s%s", result, trainTurns, terminalOutput)
			t.Error()
		}
		summary = append(summary, result)
	}

	fmt.Printf("%sSUMMARY%s\n\n", blue, reset)
	for _, result := range summary {
		fmt.Println(result)
	}
	fmt.Println("____________________")
}
