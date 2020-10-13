package gtsp

import (
	"errors"
	"fmt"
	"github.com/olegnalivajev/cmcs/pkg"
	"math"
)

type Instance struct {
	name         string
	nodeCount    int
	clusterCount int
	distances    [][]int
	clusters     map[int][]int
}

type NodeCoord struct {
	x int
	y int
}

func NewInstance(nodeCount, clusterCount int) (*Instance, error) {
	if nodeCount < clusterCount {
		return nil, errors.New("`node count` expected to be greater than `cluster count`")
	}

	// initialise distances slice
	w := make([][]int, nodeCount)
	for i := range w {
		w[i] = make([]int, nodeCount)
	}

	instance := Instance{
		nodeCount:    nodeCount,
		clusterCount: clusterCount,
		distances:    w,
		clusters:     make(map[int][]int),
	}

	instance.generateInstance()

	return &instance, nil
}

func (inst *Instance) PrintWeights() {
	for _, v := range inst.distances {
		for _, k := range v {
			fmt.Printf("%5d", k)
		}
		fmt.Println()
	}
}

func (inst *Instance) GetInstanceName() string {
	return fmt.Sprintf(`s%d-n%d-c%d`, pkg.Seed, inst.nodeCount, inst.clusterCount)
}

func (inst *Instance) GetDistance(from, to int) int {

	// since the graph isn't directional, we only save distances from smaller node
	// to higher node. vice versa has the same distance

	if from < to {
		return inst.distances[from][to]
	}
	return inst.distances[to][from]
}

func (inst *Instance) GetMinCluster() int {
	minCluster := 0
	minVertexNum := int(^uint(0) >> 1)
	for cluster, vertices := range inst.clusters {
		if len(vertices) < minVertexNum {
			minVertexNum = len(vertices)
			minCluster = cluster
		}
	}
	return minCluster
}

func (inst *Instance) VertexInCluster(v int) (int, error) {
	for cluster, vertices := range inst.clusters {
		for _, vertex := range vertices {
			if vertex == v {
				return cluster, nil
			}
		}
	}
	return 0, errors.New("no such vertex exists in any cluster")
}

func (inst *Instance) generateInstance() {

	// first create a plane with random coordinates
	// for each node

	nodes := inst.generateCoordinates()

	// now calculates the distance (distances) between
	// each node

	inst.calculateDistances(nodes)

	// generate clusters and distribute the nodes between
	// these clusters. distribution is up to the random
	// number generator
	// TODO: implement different kinds of distributions?

	inst.generateClusters()
}

func (inst *Instance) generateCoordinates() []NodeCoord {
	var nodes = make([]NodeCoord, inst.nodeCount)
	for i := 0; i < inst.nodeCount; i++ {
		coordinate := NodeCoord{
			x: pkg.GetRandomInteger(inst.nodeCount * 5),
			y: pkg.GetRandomInteger(inst.nodeCount * 5),
		}
		nodes[i] = coordinate
	}
	return nodes
}

func (inst *Instance) calculateDistances(nodes []NodeCoord) {
	for i := 0; i < inst.nodeCount; i++ {
		for j := i + 1; j < inst.nodeCount; j++ {
			inst.distances[i][j] = calculateDistance(nodes[i], nodes[j])
		}
	}
}

func (inst *Instance) generateClusters() {

	// add a single node to each of clusters to ensure
	// there's at least one node in each cluster

	for i := 0; i < inst.clusterCount; i++ {
		inst.clusters[i] = []int{i}
	}

	// distribute remaining nodes randomly between clusters

	for i := inst.clusterCount; i < inst.nodeCount; i++ {
		cluster := pkg.GetRandomInteger(inst.clusterCount)
		inst.clusters[cluster] = append(inst.clusters[cluster], i)
	}
}

// calculates Manhattan distance between 2 coordinates
func calculateDistance(c1, c2 NodeCoord) int {
	if c1.x == c2.x && c1.y == c2.y {
		return 0
	}
	return int(math.Abs(float64(c1.x-c2.x)) + math.Abs(float64(c1.y-c2.y)))
}
