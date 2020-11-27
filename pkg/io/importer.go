package io

import (
	"bufio"
	"github.com/olegnalivajev/cmcs/pkg/gtsp"
	"os"
	"strconv"
	"strings"
)

func ImportInstance(location string) (*gtsp.Instance, error) {

	// imports the instance in a format described here:
	// http://www.cs.nott.ac.uk/~pszdk/gtsp.html
	// see `Text Format (Instance)` section

	f, err := os.Open(location)
	check(err)

	// close file on exit and check for its returned error

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(f)

	// extract nodeCount

	scanner.Scan()
	nodeCountString := scanner.Text()
	nodeCountSlice := strings.Split(nodeCountString, ":")
	nodeCount, _ := strconv.Atoi(strings.TrimSpace(nodeCountSlice[len(nodeCountSlice)-1]))

	// extract clusterCount

	scanner.Scan()
	clusterCountString := scanner.Text()
	clusterCountSlice := strings.Split(clusterCountString, ":")
	clusterCount, _ := strconv.Atoi(strings.TrimSpace(clusterCountSlice[len(clusterCountSlice)-1]))

	// extract Symmetric

	scanner.Scan()
	symmetricString := scanner.Text()
	symmetricSlice := strings.Split(symmetricString, ":")
	symmetric, _ := strconv.ParseBool(strings.TrimSpace(symmetricSlice[len(symmetricSlice)-1]))

	// extract Triangle

	scanner.Scan()
	triangleString := scanner.Text()
	triangleSlice := strings.Split(triangleString, ":")
	triangle, _ := strconv.ParseBool(strings.TrimSpace(triangleSlice[len(triangleSlice)-1]))

	// extract clusters

	clusters := make(map[int][]int)

	for i := 0; i < clusterCount; i++ {
		scanner.Scan()
		clustersIRowString := scanner.Text()
		clustersIRowSlice := strings.Split(clustersIRowString, " ")

		// make an int slice

		numberOfNodesInClusterI, _ := strconv.Atoi(clustersIRowSlice[0])
		clusterI := make([]int, numberOfNodesInClusterI)
		for
		clusterI = append(clusterI, )
	}

	// extract distances

	// initialise Distances slice
	w := make([][]int, nodeCount)
	for i := range w {
		w[i] = make([]int, nodeCount)
	}

	inst := &gtsp.Instance{
		NodeCount: nodeCount,
		ClusterCount: clusterCount,
		Symmetric: symmetric,
		Triangle: triangle,
	}

	return inst, nil
}
