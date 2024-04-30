# Stations Pathfinder

This program calculates the most efficient route for trains to travel between stations on a network map.

## 1. Running the pathfinder

To run the pathfinder, use the following command:

`go run . [path to file containing network map] [start station] [end station] [number of trains]`

#### Example

```
$ go run . test_suite/fairylandNetworkMap.txt jungle desert 10
```

## 2. Information Displayed

When the program runs, it will find the most efficient route and display how the trains can be sent.

##### Example Output

The output will show the sequence of stations each train will pass through, such as:

```
T1-grasslands T2-farms T3-green_belt
T1-suburbs T2-downtown T3-village T4-grasslands T5-farms T6-green_belt
T1-clouds T2-metropolis T3-mountain T4-suburbs T5-downtown T6-village T7-grasslands T8-farms T9-green_belt
T1-wetlands T2-industrial T3-treetop T4-clouds T5-metropolis T6-mountain T7-suburbs T8-downtown T9-village T10-grasslands
T1-desert T2-desert T3-desert T4-wetlands T5-industrial T6-treetop T7-clouds T8-metropolis T9-mountain T10-suburbs
T4-desert T5-desert T6-desert T7-wetlands T8-industrial T9-treetop T10-clouds
T7-desert T8-desert T9-desert T10-wetlands
T10-desert
```

## 3. Checking Total Turns

To check the total number of turns used for the fastest route, You can use the following command:  
`go run . [path to file containing network map] [start station] [end station] [number of trains] | wc -l`

#### Example

```
$ go run . test_suite/fairylandNetworkMap.txt jungle desert 10 | wc -l
8
```

## 4. Running The Suite Of Tests

The suite of tests is created for testing the program with all of the example inputs.  
Run the suite of tests by following these instructions:

1. Compile the code by using command `go build`
2. The binary file named `trains` is now located in the working directory
3. Run the suite of tests by using the command `go test`
4. Test results are now visible in the terminal

#### Example output

```
PASSED Food Network Map: 4 trains between bond_square and space_port
Number of train turns used: 6
Maximum train turns usable: 6

Train turns:
T1-apple_avenue
T1-orange_junction T2-apple_avenue
T1-space_port T2-orange_junction T3-apple_avenue
T2-space_port T3-orange_junction T4-apple_avenue
T3-space_port T4-orange_junction
T4-space_port
```

Result for each test can be:

- PASSED
- FAILED (too many train turns used)
- INVALID INPUT (program reports an error regarding input)

The suite of tests output ends in a summary where you can easily see which test passed and which failed.

## 5. Algorithm Operation Principles

1. Find all routes from the start station to the end station
2. Find all possible combinations of routes with unique stations
3. Find fastest combination to use for number of trains specified by the user
4. Distribute the right number of trains to each route used and print train turns to terminal
