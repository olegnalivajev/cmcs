package main

import (
	"fmt"
	"github.com/olegnalivajev/cmcs/cmd"
	"github.com/olegnalivajev/cmcs/pkg/gtsp"
)

func main() {
	cmd.Execute()
	test()
}

func test() {
	inst, _ := gtsp.NewInstance(22, 5)
	fmt.Println(inst.GetDistance(0, 3))
	fmt.Println(inst.GetDistance(3, 0))
}

