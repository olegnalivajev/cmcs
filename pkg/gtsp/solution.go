package gtsp

import (
	"github.com/olegnalivajev/cmcs/pkg"
)

type Solution struct {
	Instance    Instance
	Distance    int
	Vertices    []int // vertex in cluster(index)
	PrevCluster []int // cluster (value) preceding cluster (index)
	NextCluster []int // cluster (value) succeeding cluster (index)
}

func GenerateSolution(instance Instance) *Solution {
	solution := Solution{
		Instance:    instance,
		Distance:    0,
		Vertices:    make([]int, instance.clusterCount),
		PrevCluster: make([]int, instance.clusterCount),
		NextCluster: make([]int, instance.clusterCount),
	}
	solution.generateInitialSolution()
	return &solution
}

func (s *Solution) UpdateDistance(amount int) {
	s.Distance += amount
}

// TODO: implement
func (s *Solution) CalculateDistance() int {
	return 0
}

// TODO: implement
func (s *Solution) IsFeasible() bool {

	return true
}

// TODO: implement
func (s *Solution) deepCopy() *Solution {
	return &Solution{

	}
}

func (s *Solution) generateInitialSolution() {

	clusters := make([]int, s.Instance.clusterCount-1)

	// first, generate an array of clusters excluding cluster 0, as our generator starts from it anyway

	for i := 0; i < len(clusters); i++ {
		clusters[i] = i+1
	}

	// pick a cluster at random that would be a next cluster to cluster 0.
	// remove it from available clusters, however it's now recorded as current cluster
	// so we can iteratively pick random clusters to follow the current one

	rnd := pkg.GetRandomInteger(len(clusters))
	curr := clusters[rnd]
	s.NextCluster[0] = curr
	s.PrevCluster[curr] = 0
	clusters = remove(clusters, rnd)

	// iteratively pick a next cluster that will follow the current one.
	// since we have already found a cluster succeeding cluster 0, and cluster 0
	// will also follow the last cluster from the slice (our graph has a single cycle,
	// starting at node `x`, going through each node, and returning into node `x`),
	// the iteration starts from i = 2

	for i := 2; i < s.Instance.clusterCount; i++ {
		rnd = pkg.GetRandomInteger(len(clusters))
		cluster := clusters[rnd]
		s.NextCluster[curr] = cluster
		s.PrevCluster[cluster] = curr
		curr = cluster
		clusters = remove(clusters, rnd)
	}

	// at this point `curr` cluster is the last cluster we processed, so cluster 0
	// will follow it, therefore the preceding cluster to cluster 0 is `curr` cluster

	s.PrevCluster[0] = curr


	// select a random node from each cluster

	for i := 0; i < s.Instance.clusterCount; i++ {
		rndIndex := len(s.Instance.clusters[i])
		s.Vertices[i] = s.Instance.clusters[i][pkg.GetRandomInteger(rndIndex)]
	}
}

// removes an int element from a slice with no duplicates
func remove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}