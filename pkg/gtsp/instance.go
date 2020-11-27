package gtsp

import (
	"errors"
	"fmt"
	"github.com/olegnalivajev/cmcs/pkg"
	"math"
)

type Instance struct {
	Triangle     bool
	Symmetric    bool
	NodeCount    int
	ClusterCount int
	Distances    [][]int
	Clusters     map[int][]int
}

type NodeCoord struct {
	x int
	y int
}

func NewInstance(nodeCount, clusterCount int) (*Instance, error) {
	if nodeCount < clusterCount {
		return nil, errors.New("`node count` expected to be greater than `cluster count`")
	}

	// initialise Distances slice
	w := make([][]int, nodeCount)
	for i := range w {
		w[i] = make([]int, nodeCount)
	}

	instance := Instance{
		Triangle: false,
		Symmetric: true,
		NodeCount:    nodeCount,
		ClusterCount: clusterCount,
		Distances:    w,
		Clusters:     make(map[int][]int),
	}

	instance.generateInstance()

	return &instance, nil
}

func (inst *Instance) PrintWeights() {
	for _, v := range inst.Distances {
		for _, k := range v {
			fmt.Printf("%5d", k)
		}
		fmt.Println()
	}
}

func (inst *Instance) GetInstanceName() string {
	return fmt.Sprintf(`s%d-n%d-c%d`, pkg.Seed, inst.NodeCount, inst.ClusterCount)
}

func (inst *Instance) GetDistance(from, to int) int {

	// since the graph isn't directional, we only save Distances from smaller node
	// to higher node. vice versa has the same distance

	if from < to {
		return inst.Distances[from][to]
	}
	return inst.Distances[to][from]
}

func (inst *Instance) GetMinCluster() int {
	minCluster := 0
	minVertexNum := int(^uint(0) >> 1)
	for cluster, vertices := range inst.Clusters {
		if len(vertices) < minVertexNum {
			minVertexNum = len(vertices)
			minCluster = cluster
		}
	}
	return minCluster
}

func (inst *Instance) VertexInCluster(v int) (int, error) {

	// returns a clusters in which given vertex is placed/
	// if no such vertex exists returns an error

	for cluster, vertices := range inst.Clusters {
		for _, vertex := range vertices {
			if vertex == v {
				return cluster, nil
			}
		}
	}
	return 0, errors.New("no such vertex exists in any cluster")
}

func (inst *Instance) DeepCopy() *Instance {

	// slice is a reference type, therefore we have to
	// iterate over the og slice and copy values one by one

	dist := make([][]int, inst.NodeCount)
	for i := range dist {
		dist[i] = make([]int, inst.NodeCount)
		for j := range dist[i] {
			dist[i][j] = inst.Distances[i][j]
		}
	}

	// same applies to map of Clusters

	cls := make(map[int][]int)
	for k, v := range inst.Clusters {
		nodes := make([]int, len(v))
		for i, _ := range v {
			nodes[i] = v[i]
		}
		cls[k] = nodes
	}

	return &Instance{
		NodeCount:    inst.NodeCount,
		ClusterCount: inst.ClusterCount,
		Distances:    dist,
		Clusters:     cls,
	}
}

func (inst *Instance) generateInstance() {

	// first create a plane with random coordinates
	// for each node

	nodes := inst.generateCoordinates()

	// now calculates the distance (Distances) between
	// each node

	inst.calculateDistances(nodes)

	// generate Clusters and distribute the nodes between
	// these Clusters. distribution is up to the random
	// number generator
	// TODO: implement different kinds of distributions?

	inst.generateClusters()
}

func (inst *Instance) generateCoordinates() []NodeCoord {
	var nodes = make([]NodeCoord, inst.NodeCount)
	for i := 0; i < inst.NodeCount; i++ {
		coordinate := NodeCoord{
			x: pkg.GetRandomInteger(inst.NodeCount * 5),
			y: pkg.GetRandomInteger(inst.NodeCount * 5),
		}
		nodes[i] = coordinate
	}
	return nodes
}

func (inst *Instance) calculateDistances(nodes []NodeCoord) {
	for i := 0; i < inst.NodeCount; i++ {
		for j := i + 1; j < inst.NodeCount; j++ {
			distance := calculateDistance(nodes[i], nodes[j])
			inst.Distances[i][j] = distance
			inst.Distances[j][i] = distance
		}
	}
}

func (inst *Instance) generateClusters() {

	// add a single node to each cluster to ensure
	// there's at least one node

	for i := 0; i < inst.ClusterCount; i++ {
		inst.Clusters[i] = []int{i}
	}

	// distribute remaining nodes randomly between Clusters

	for i := inst.ClusterCount; i < inst.NodeCount; i++ {
		cluster := pkg.GetRandomInteger(inst.ClusterCount)
		inst.Clusters[cluster] = append(inst.Clusters[cluster], i)
	}
}

// calculates Manhattan distance between 2 coordinates
func calculateDistance(c1, c2 NodeCoord) int {
	if c1.x == c2.x && c1.y == c2.y {
		return 0
	}
	return int(math.Abs(float64(c1.x-c2.x)) + math.Abs(float64(c1.y-c2.y)))
}
