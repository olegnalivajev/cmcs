package main

import (
	"fmt"
	"github.com/olegnalivajev/cmcs/pkg/io"
)

func main() {
	//instance, err := gtsp.NewInstance(40, 7)
	//if err != nil {
	//	panic(err)
	//}
	//io.ExportInstance(*instance, "C:\\Users\\Oleg\\Documents\\Projects\\cmcs")

	inst, _ := io.ImportInstance("C:\\Users\\Oleg\\Documents\\Projects\\cmcs\\test_instance.txt")
	fmt.Println(inst)
}