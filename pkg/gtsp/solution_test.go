package gtsp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution_InsertCluster(t *testing.T) {
	inst, err := NewInstance(5, 5)
	assert.True(t, err == nil)

	inst.Clusters = map[int][]int{
		0: {0},
		1: {1},
		2: {2},
		3: {3},
		4: {4},
	}

	inst.Distances = [][]int{
		{0, 8, 9, 10, 11},
		{0, 0, 12, 13, 14},
		{0, 0, 0, 15, 16},
		{0, 0, 0, 0, 17},
		{0, 0, 0, 0, 0},
	}

	solution := GenerateSolution(*inst)

	// hand-defined cluster sequence: 0 -> 1 -> 2 -> 3 -> 4 -> 0

	solution.PrevCluster = []int{4, 0, 1, 2, 3}
	solution.NextCluster = []int{1, 2, 3, 4, 0}

	// recalculate the distance

	solution.CalculateDistance()

	// since we have defined distance ourselves, we can assert if it's what we expect
	// just as a sanity check

	// 0-1 : 8
	// 1-2 : 12
	// 2-3 : 15
	// 3-4 : 17
	// 4-0 : 11
	//     = 63

	assert.Equal(t, 63, solution.Distance)

	// let's insert cluster 0 after cluster 3, so the structure of the graph will change to:
	// 0 -> 4 -> 1 -> 2 -> 3 -> 0

	solution.InsertCluster(0, 3)

	expectedPrev := []int{3, 4, 1, 2, 0}
	expectedNext := []int{4, 2, 3, 0, 1}

	// removed edges:
	// 0-1 : 8
	// 4-0 : 11
	// 3-4 : 17
	// added edges:
	// 4-1 : 14
	// 3-0 : 10
	// 0-4 : 11

	expectedDistance := 62

	assert.Equal(t, solution.NextCluster, expectedNext)
	assert.Equal(t, solution.PrevCluster, expectedPrev)
	assert.Equal(t, solution.Distance, expectedDistance)
}

func TestSolution_SwapVertexInCluster_ClusterSizeOne(t *testing.T) {
	inst, err := NewInstance(10, 3)
	assert.True(t, err == nil)

	// only a single vertex in cluster 1

	inst.Clusters[0] = []int{0}

	solution := GenerateSolution(*inst)

	vertex := solution.Vertices[0]

	solution.SwapVertexInCluster(0)

	// the vertex should have not changed

	assert.Equal(t, solution.Vertices[0], vertex)
}

func TestSolution_SwapVertexInCluster(t *testing.T) {
	inst, err := NewInstance(10, 3)
	assert.True(t, err == nil)

	inst.Clusters[0] = []int{0, 3}

	solution := GenerateSolution(*inst)

	initialVertex := solution.Vertices[0]

	solution.SwapVertexInCluster(0)

	assert.NotEqual(t, solution.Vertices[0], initialVertex)
}

func TestSolution_UpdateDistance(t *testing.T) {
	inst, err := NewInstance(10, 3)
	assert.True(t, err == nil)

	solution := GenerateSolution(*inst)

	distance := solution.Distance
	incrementBy := 5

	solution.UpdateDistance(incrementBy)

	assert.Equal(t, solution.Distance, distance+incrementBy)
}

func TestSolution_CalculateDistance(t *testing.T) {
	inst, err := NewInstance(5, 3)
	assert.True(t, err == nil)

	inst.Distances = [][]int{
		{0, 5, 8, 12, 10},
		{0, 0, 10, 4, 13},
		{0, 0, 0, 12, 15},
		{0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0},
	}

	inst.Clusters = map[int][]int{
		0: {0},
		1: {1, 3},
		2: {2, 4},
	}

	solution := GenerateSolution(*inst)
	solution.Vertices = []int{0, 1, 4}

	// recalculate the distance

	solution.CalculateDistance()

	// expected distance =
	// 0 to 1 = 5 +
	// 1 to 4 = 13 +
	// 4 to 0 = 10
	// = 28

	assert.Equal(t, 28, solution.Distance)
}

func TestSolution_IsFeasible_IncorrectVertex(t *testing.T) {
	inst, err := NewInstance(5, 3)
	assert.True(t, err == nil)

	solution := GenerateSolution(*inst)

	// assign vertex from a different cluster
	// e.g. first vertex from cluster 1 as a vertex in cluster 0

	solution.Vertices[0] = inst.Clusters[1][0]

	assert.False(t, solution.IsFeasible())
}

func TestSolution_IsFeasible_PreviousNextDoNotCorrespond(t *testing.T) {
	inst, err := NewInstance(5, 3)
	assert.True(t, err == nil)

	solution := GenerateSolution(*inst)

	// assign vertex from a different cluster
	// e.g. first vertex from cluster 1 as a vertex in cluster 0

	solution.NextCluster = []int{1, 2, 0}
	solution.PrevCluster = []int{2, 1, 0}

	assert.False(t, solution.IsFeasible())
}

func TestSolution_GenerateSolution(t *testing.T) {

	// checks if generated solution is feasible

	inst, err := NewInstance(5, 3)
	assert.True(t, err == nil)

	solution := GenerateSolution(*inst)

	assert.True(t, solution.IsFeasible())
}

func TestSolution_DeepCopy(t *testing.T) {
	inst, err := NewInstance(110, 5)
	assert.True(t, err == nil)

	solution := GenerateSolution(*inst)
	deepCopy := solution.DeepCopy()

	assert.Equal(t, *solution, *deepCopy)
	assert.False(t, &solution.Vertices == &deepCopy.Vertices)
	assert.False(t, &solution.PrevCluster == &deepCopy.PrevCluster)
	assert.False(t, &solution.NextCluster == &deepCopy.NextCluster)
	assert.False(t, &solution.Distance == &deepCopy.Distance)
	assert.False(t, &solution.Instance == &deepCopy.Instance)
}
