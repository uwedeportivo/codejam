// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/7214486/dashboard#s=p2

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/uwedeportivo/codejam"
)

type data struct {
	testIndex int
	vertices  []string
	edges     [][]int
}

type node struct {
	vertex int
	color  string
	kids   []*node
	tr     string
}

type byTrace []*node

func (as byTrace) Len() int      { return len(as) }
func (as byTrace) Swap(i, j int) { as[i], as[j] = as[j], as[i] }

func (as byTrace) Less(i, j int) bool {
	return as[i].trace() < as[j].trace()
}

func (t *node) trace() string {
	if len(t.tr) > 0 {
		return t.tr
	}

	var buf bytes.Buffer

	if len(t.kids) > 0 {
		buf.WriteString("(")
	}

	buf.WriteString(t.color)

	sort.Sort(byTrace(t.kids))

	for _, tk := range t.kids {
		buf.WriteString(" ")
		buf.WriteString(tk.trace())
	}

	if len(t.kids) > 0 {
		buf.WriteString(")")
	}

	t.tr = buf.String()
	return t.tr
}

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	d := &data{
		testIndex: testCaseIndex,
	}

	n := pr.ReadInt()
	d.vertices = make([]string, n)
	d.edges = make([][]int, n)

	for i := 0; i < n; i++ {
		d.vertices[i] = pr.ReadString()
	}

	for i := 0; i < n-1; i++ {
		edge := pr.ReadInts(nil)
		a, b := edge[0]-1, edge[1]-1

		d.edges[a] = append(d.edges[a], b)
		d.edges[b] = append(d.edges[b], a)
	}

	return d
}

func removeFromList(xs []int, v int) []int {
	for k, x := range xs {
		if x == v {
			if k == len(xs)-1 {
				return xs[:k]
			} else {
				return append(xs[:k], xs[k+1:]...)
			}
		}
	}
	panic("value not in list")
}

func addUniqueToList(xs []int, v int) []int {
	found := false
	for _, x := range xs {
		if v == x {
			found = true
			break
		}
	}
	if !found {
		return append(xs, v)
	}
	return xs
}

func buildTree(d *data) *node {
	v2n := make(map[int]*node)
	var work []int

	for i := 0; i < len(d.vertices); i++ {
		if len(d.edges[i]) == 1 {
			work = append(work, i)
		}
	}

	for {
		var parents []int

		for _, k := range work {
			es := d.edges[k]

			if len(es) == 1 {
				e := es[0]

				parent, ok := v2n[e]
				if !ok {
					parent = &node{
						vertex: e,
						color:  d.vertices[e],
					}
					v2n[e] = parent
				}
				parents = addUniqueToList(parents, e)

				child, ok := v2n[k]
				if !ok {
					child = &node{
						vertex: k,
						color:  d.vertices[k],
					}
					v2n[k] = child
				}
				parent.kids = append(parent.kids, child)
				d.edges[k] = nil
				d.edges[e] = removeFromList(d.edges[e], k)
			}
		}

		if len(parents) == 1 {
			return v2n[parents[0]]
		}

		work = nil
		for _, k := range parents {
			if len(d.edges[k]) == 1 {
				work = append(work, k)
			}
		}
	}
	panic("no tree root")
}

func splitRoot(t *node, index int) (*node, *node) {
	ckids := make([]*node, len(t.kids)-1)
	copy(ckids[:index], t.kids[:index])
	copy(ckids[index:], t.kids[index+1:])

	ct := &node{
		vertex: t.vertex,
		color:  t.color,
		kids:   ckids,
	}
	return ct, t.kids[index]
}

func symmetric(d *data, t *node, verticals int) bool {
	if verticals == 2 && len(d.vertices)%2 == 0 {
		// special case with nobody on the symmetry axis
		for i := 0; i < len(t.kids); i++ {
			a, b := splitRoot(t, i)

			if a.trace() == b.trace() {
				return true
			}
		}
	}

	traceHisto := make(map[string]int)
	trace2node := make(map[string]*node)

	for _, tk := range t.kids {
		ts := tk.trace()
		traceHisto[ts] = traceHisto[ts] + 1
		trace2node[ts] = tk
	}

	var oddTraces []string

	for ts, nr := range traceHisto {
		if nr%2 == 1 {
			oddTraces = append(oddTraces, ts)
		}
	}

	if len(oddTraces) > verticals {
		return false
	}

	for _, ts := range oddTraces {
		tk := trace2node[ts]

		if !symmetric(d, tk, 1) {
			return false
		}
	}
	return true
}

func solve(pr *codejam.Problem, d *data) {
	root := buildTree(d)

	if symmetric(d, root, 2) {
		pr.Write(fmt.Sprintf("Case #%d: SYMMETRIC\n", d.testIndex))
	} else {
		pr.Write(fmt.Sprintf("Case #%d: NOT SYMMETRIC\n", d.testIndex))
	}
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
