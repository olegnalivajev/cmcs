package gtsp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	inst.distances = [][]int{
		{0, 5, 8, 12, 10},
		{0, 0, 10, 4, 13},
		{0, 0, 0, 12, 15},
		{0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0},
	}

	inst.clusters = map[int][]int{
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

	solution.Vertices[0] = inst.clusters[1][0]

	assert.False(t, solution.IsFeasible())
}

func TestSolution_IsFeasible_PreviousNextDoNotCorrespond(t *testing.T) {
	inst, err := NewInstance(5, 3)
	assert.True(t, err == nil)

	solution := GenerateSolution(*inst)

	// assign vertex from a different cluster
	// e.g. first vertex from cluster 1 as a vertex in cluster 0

	solution.NextCluster = []int{1,2,0}
	solution.PrevCluster = []int{2,1,0}

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
