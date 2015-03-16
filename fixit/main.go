// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/635101/dashboard#s=p0

package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/uwedeportivo/codejam"
)

type kids []*node

type node struct {
	val string
	ks  kids
}

func (nd *node) declarePath(path string) int {
	xs := strings.Split(path, "/")

	return nd.declare(xs[1:], 0)
}

func (nd *node) declare(path []string, cursor int) int {
	if cursor == len(path) {
		return 0
	}

	j := sort.Search(len(nd.ks), func(i int) bool {
		return nd.ks[i].val >= path[cursor]
	})

	var k *node
	inserts := 0

	if j < len(nd.ks) && nd.ks[j].val == path[cursor] {
		k = nd.ks[j]
	} else {
		k = &node{
			val: path[cursor],
		}
		inserts = 1

		nd.ks = append(nd.ks, nil)
		copy(nd.ks[j+1:], nd.ks[j:])
		nd.ks[j] = k
	}

	return k.declare(path, cursor+1) + inserts
}

type data struct {
	testIndex int
	es        []string
	ns        []string
}

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	d := &data{
		testIndex: testCaseIndex,
	}

	nm := pr.ReadInts(nil)
	n := nm[0]
	m := nm[1]

	d.es = make([]string, n)
	d.ns = make([]string, m)

	for i := 0; i < n; i++ {
		d.es[i] = pr.ReadString()
	}

	for i := 0; i < m; i++ {
		d.ns[i] = pr.ReadString()
	}

	return d
}

func solve(pr *codejam.Problem, d *data) {
	var nw int

	root := &node{}

	for _, path := range d.es {
		root.declarePath(path)
	}

	for _, path := range d.ns {
		nw += root.declarePath(path)
	}

	pr.Write(fmt.Sprintf("Case #%d: %d\n", d.testIndex, nw))
}

func main() {
	help := flag.Bool("help", false, "show this message")
	inFile := flag.String("in", "", "input filename (required)")
	outFile := flag.String("out", "", "output filename (stdout if omitted)")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if len(*inFile) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	pr := codejam.NewProblem(*inFile, *outFile)

	numTestCases := pr.ReadInt()

	if numTestCases < 1 {
		panic(fmt.Errorf("no testcases available"))
	}

	for testIndex := 1; testIndex <= numTestCases; testIndex++ {
		d := parse(pr, testIndex)
		solve(pr, d)
	}

	pr.Close()
}
