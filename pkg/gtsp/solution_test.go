package gtsp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateInitialSolution(t *testing.T) {
	inst, err := NewInstance(110, 8)

	assert.True(t, err == nil)
	s := GenerateSolution(*inst)
	fmt.Println(s.PrevCluster)
	fmt.Println(s.NextCluster)
}