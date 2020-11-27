package io

import (
	"bufio"
	"fmt"
	"github.com/olegnalivajev/cmcs/pkg/gtsp"
	"os"
	"strconv"
)

func ExportInstance(instance gtsp.Instance, location string) {

	// exports the instance in a format described here:
	// http://www.cs.nott.ac.uk/~pszdk/gtsp.html
	// see `Text Format (Instance)` section

	f, err := os.Create(location+"/"+instance.GetInstanceName()+".txt")
	check(err)

	// close file on exit and check for its returned error

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// make a write buffer

	w := bufio.NewWriter(f)

	// headers

	_, err = w.WriteString("N: " + strconv.Itoa(instance.NodeCount) + "\n")
	check(err)
	_, err = w.WriteString("M: " + strconv.Itoa(instance.ClusterCount) + "\n")
	check(err)
	_, err = w.WriteString("Symmetric: " + strconv.FormatBool(instance.Symmetric) +"\n")
	check(err)
	_, err = w.WriteString("Triangle: " + strconv.FormatBool(instance.Triangle) + "\n")
	check(err)

	// clusters

	for _, nodes := range instance.Clusters {
		_, err = w.WriteString(fmt.Sprintf("%d ", len(nodes)))
		check(err)
		for _, node := range nodes {
			_, err = w.WriteString(fmt.Sprintf("%d ", node))
			check(err)
		}
		_, err := w.WriteString("\n")
		check(err)
	}

	// distance matrix

	for _, rows := range instance.Distances {
		for _, dist := range rows {
			_, err := w.WriteString(fmt.Sprintf("%d ", dist))
			check(err)
		}
		_, err := w.WriteString("\n")
		check(err)
	}

	err = w.Flush()
	check(err)
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
