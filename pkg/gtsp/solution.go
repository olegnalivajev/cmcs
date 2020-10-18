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
	solution.CalculateDistance()
	return &solution
}

func (s *Solution) UpdateDistance(amount int) {
	s.Distance += amount
}

func (s *Solution) CalculateDistance() int {
	s.Distance = 0
	for i, v := range s.Vertices {
		s.Distance += s.Instance.GetDistance(v, s.Vertices[s.NextCluster[i]])
	}
	return 0
}

func (s *Solution) InsertCluster(cluster, afterCluster int) {

	// selected cluster, x, is extracted from it's position and inserted
	// after cluster z. i.e. prev[x] no longer points to x, but to next[x]
	// likewise, z no longer points to next[z], but to x, and prev[next[z]]
	// no longer points to z but to x

	// first update pointers to previous and next clusters

	before := s.PrevCluster[cluster]
	after := s.NextCluster[cluster]
	between := s.NextCluster[afterCluster]

	s.NextCluster[before] = after
	s.NextCluster[afterCluster] = cluster
	s.NextCluster[cluster] = between
	s.PrevCluster[after] = before
	s.PrevCluster[cluster] = afterCluster
	s.PrevCluster[between] = cluster

	// then recalculate the distance. to make it more efficient and avoid new traversal
	// we subtract the weights of removed edges, and add the weights of new edges

	clusterVertex := s.Vertices[cluster]
	beforeVertex := s.Vertices[before]
	afterVertex := s.Vertices[after]
	newBeforeVertex := s.Vertices[afterCluster]
	newAfterVertex := s.Vertices[between]

	s.Distance -= s.Instance.GetDistance(beforeVertex, clusterVertex)
	s.Distance -= s.Instance.GetDistance(clusterVertex, afterVertex)
	s.Distance -= s.Instance.GetDistance(newBeforeVertex, newAfterVertex)
	s.Distance += s.Instance.GetDistance(beforeVertex, afterVertex)
	s.Distance += s.Instance.GetDistance(newBeforeVertex, clusterVertex)
	s.Distance += s.Instance.GetDistance(clusterVertex, newAfterVertex)
}

func (s *Solution) SwapVertexInCluster(cluster int) {

	// swaps current vertex in cluster to another random vertex from the same cluster.
	// the swap is guaranteed, unless cluster is of size 1

	if len(s.Instance.clusters[cluster]) == 1 {
		return
	}

	rndIndex := pkg.GetRandomInteger(len(s.Instance.clusters[cluster]))
	newVertex := s.Instance.clusters[cluster][rndIndex]

	// if we have selected the same vertex, we gonna call a function again,
	// until a different vertex is selected

	if newVertex == s.Vertices[cluster] {
		s.SwapVertexInCluster(cluster)
		return
	}

	s.Vertices[cluster] = newVertex
}

func (s *Solution) IsFeasible() bool {

	// check if vertex actually exists in the corresponding cluster

	for i, v := range s.Vertices {
		found := false
		for _, vertex := range s.Instance.clusters[i] {
			if v == vertex {
				found = true
			}
		}
		if !found {
			return false
		}
	}

	// check if slices with next & previous clusters correspond

	for i := 0; i < s.Instance.clusterCount; i++ {
		if s.PrevCluster[s.NextCluster[i]] != i {
			return false
		}
	}

	// traverse the entire graph and check if the graph is cyclic
	// i.e. when it starts at 0, it should end at 0 as well

	firstCluster := 0
	for i := 0; i < s.Instance.clusterCount; i++ {
		firstCluster = s.NextCluster[firstCluster]
	}
	return firstCluster == 0

}

func (s *Solution) DeepCopy() *Solution {

	vrt := make([]int, len(s.Vertices))
	copy(vrt, s.Vertices)

	prev := make([]int, len(s.PrevCluster))
	copy(prev, s.PrevCluster)

	next := make([]int, len(s.NextCluster))
	copy(next, s.NextCluster)

	return &Solution{
		Instance:    *s.Instance.DeepCopy(),
		Distance:    s.Distance,
		Vertices:    vrt,
		PrevCluster: prev,
		NextCluster: next,
	}

}

func (s *Solution) generateInitialSolution() {

	clusters := make([]int, s.Instance.clusterCount-1)

	// first, generate an array of clusters excluding cluster 0, as our generator starts from it anyway

	for i := 0; i < len(clusters); i++ {
		clusters[i] = i + 1
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
