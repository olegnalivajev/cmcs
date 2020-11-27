package gtsp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInstanceClusterCountGreaterThanNodeCount(t *testing.T) {
	_, err := NewInstance(3, 10)
	assert.Equal(t, "`node count` expected to be greater than `cluster count`", err.Error())
}

func TestNewInstanceHappyPath(t *testing.T) {
	instance, err := NewInstance(10, 3)
	assert.True(t, err == nil)
	assert.EqualValues(t, 10, instance.NodeCount)
	assert.EqualValues(t, 3, instance.ClusterCount)
	assert.True(t, len(instance.Clusters) == instance.ClusterCount)
}

func TestInstance_GetInstanceName(t *testing.T) {
	instance, err := NewInstance(10, 3)
	assert.True(t, err == nil)
	assert.EqualValues(t, "s1-n10-c3", instance.GetInstanceName())
}

func TestInstance_GetDistance(t *testing.T) {
	weights := [][]int{
		{0, 2, 3},
		{0, 0, 4},
		{0, 0, 0},
	}

	instance, err := NewInstance(10, 3)
	assert.True(t, err == nil)

	instance.Distances = weights

	assert.EqualValues(t, 4, instance.GetDistance(1, 2))
	assert.EqualValues(t, 4, instance.GetDistance(2, 1))
}

func TestInstance_GetMinCluster(t *testing.T) {
	clusters := map[int][]int{
		0: {1, 2, 3},
		1: {4},
		2: {5, 6, 7, 8},
	}

	instance, err := NewInstance(10, 3)
	assert.True(t, err == nil)

	instance.Clusters = clusters
	assert.EqualValues(t, 1, instance.GetMinCluster())
}

func TestInstance_VertexInCluster_VertexDoesNotExists(t *testing.T) {
	instance, err := NewInstance(10, 3)
	assert.True(t, err == nil)

	_, err = instance.VertexInCluster(22)
	assert.True(t, err != nil)
	assert.EqualValues(t, "no such vertex exists in any cluster", err.Error())
}

func TestInstance_DeepCopy(t *testing.T) {
	instance, err := NewInstance(10, 3)
	assert.True(t, err == nil)

	deepCopy := instance.DeepCopy()

	assert.Equal(t, *instance, *deepCopy)
	assert.False(t, &instance == &deepCopy)
	assert.False(t, &instance.Clusters == &deepCopy.Clusters)
	assert.False(t, &instance.Distances == &deepCopy.Distances)
	assert.False(t, &instance.NodeCount == &deepCopy.NodeCount)
	assert.False(t, &instance.ClusterCount == &deepCopy.ClusterCount)
}

func TestInstance_VertexInCluster(t *testing.T) {
	clusters := map[int][]int{
		0: {1, 2, 3},
		1: {4},
		2: {5, 6, 7, 8},
	}

	instance, err := NewInstance(10, 3)
	assert.True(t, err == nil)

	instance.Clusters = clusters

	cluster, err := instance.VertexInCluster(4)

	assert.True(t, err == nil)
	assert.EqualValues(t, 1, cluster)
}

func TestInstance_CalculateDistances(t *testing.T) {
	instance, err := NewInstance(3, 1)
	assert.True(t, err == nil)

	nodes := []NodeCoord{
		{4, 8},
		{3, 6},
		{2, 9},
	}

	instance.calculateDistances(nodes)

	expectedDistances := [][]int{
		{0, 3, 3},
		{0, 0, 4},
		{0, 0, 0},
	}

	assert.EqualValues(t, instance.Distances, expectedDistances)
}

func Test_CalculateDistance(t *testing.T) {
	node1 := NodeCoord{3, 5}
	node2 := NodeCoord{6, 8}

	expectedDistance := 6

	assert.EqualValues(t, expectedDistance, calculateDistance(node1, node2))
}

func TestInstance_GenerateClustersHasAtLeastASingleNodeInCluster(t *testing.T) {
	instance, err := NewInstance(3, 2)
	assert.True(t, err == nil)
	assert.True(t, len(instance.Clusters[0]) > 0)
	assert.True(t, len(instance.Clusters[1]) > 0)
}
