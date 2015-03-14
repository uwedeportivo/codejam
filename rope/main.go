// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/619102/dashboard#s=p0

package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/uwedeportivo/codejam"
	"github.com/willf/bitset"
)

type pair [2]int

type data struct {
	testIndex int
	pairs     []pair
}

type pairSorter struct {
	indices []int
	pairs   []pair
	column  int
}

func (ps *pairSorter) Len() int {
	return len(ps.pairs)
}

func (ps *pairSorter) Less(i, j int) bool {
	return ps.pairs[ps.indices[i]][ps.column] < ps.pairs[ps.indices[j]][ps.column]
}

func (ps *pairSorter) Swap(i, j int) {
	ps.indices[i], ps.indices[j] = ps.indices[j], ps.indices[i]
}

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	d := new(data)
	d.testIndex = testCaseIndex

	n := pr.ReadInt()
	d.pairs = make([]pair, n)

	for i := 0; i < n; i++ {
		xs := pr.ReadInts(nil)
		d.pairs[i][0] = xs[0]
		d.pairs[i][1] = xs[1]
	}
	return d
}

func solve(pr *codejam.Problem, d *data) {
	n := len(d.pairs)

	ais := make([]int, n)
	bis := make([]int, n)

	for i := 0; i < n; i++ {
		ais[i] = i
		bis[i] = i
	}

	psa := &pairSorter{
		indices: ais,
		pairs:   d.pairs,
		column:  0,
	}

	psb := &pairSorter{
		indices: bis,
		pairs:   d.pairs,
		column:  1,
	}

	sort.Sort(psa)
	sort.Sort(psb)

	var nw int

	for i := 0; i < n-1; i++ {
		v := d.pairs[ais[i]][1]

		k := sort.Search(n, func(j int) bool {
			return d.pairs[bis[j]][1] >= v
		})

		// count common indices from ais[i+1:] and bis[:k]
		var ba bitset.BitSet
		var bb bitset.BitSet

		for _, c := range ais[i+1:] {
			ba.Set(uint(c))
		}

		for _, c := range bis[:k] {
			bb.Set(uint(c))
		}

		nw += int(ba.IntersectionCardinality(&bb))
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
