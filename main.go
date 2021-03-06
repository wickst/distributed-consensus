package main

import (
	"github.com/wickst/distributed-consensus/util"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// see if there is a commandline package to make named command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments given. Please use a valid argument for the number of nodes.")
		os.Exit(1)
	}
	args := os.Args[1:]

	nodeNumber, err := strconv.Atoi(args[0])
	// TODO: check type of argument passed in via commandline
	if err != nil {
		fmt.Println("Incorrect arguments given. Please use the following command ./")
		os.Exit(1)
	}

	// need to name variables/files betterer
	var graph graph.Graph

	graph.CreateRandomGraph(nodeNumber)

	graph.BroadcastNodeInfo()

	graph.ConsensusResult()
}
