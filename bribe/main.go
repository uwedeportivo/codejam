// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/189252/dashboard#s=p2

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/uwedeportivo/codejam"
)

type data struct {
	testIndex int
	nc        int
	released  []int
}

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	d := &data{
		testIndex: testCaseIndex,
	}

	pq := pr.ReadInts(nil)

	d.nc = pq[0]

	d.released = make([]int, pq[1])
	pr.ReadInts(d.released)

	return d
}

// invariant: prs[l:r] are all still to be released
// a is first empty cell on the left of prs
// b is first empty cell on the right of prs
func coins(prs []int, v [4]int, memo map[[4]int]int) int {
	a, b, l, r := v[0], v[1], v[2], v[3]

	n := r - l
	if n == 0 {
		return 0
	}
	if n == 1 {
		return b - a - 2
	}

	if c, ok := memo[v]; ok {
		return c
	}

	min := coins(prs, [4]int{prs[l], b, l + 1, r}, memo) + b - a - 2
	for i := l + 1; i < r; i++ {
		c := coins(prs, [4]int{a, prs[i], l, i}, memo) + coins(prs, [4]int{prs[i], b, i + 1, r}, memo) + b - a - 2
		if c < min {
			min = c
		}
	}
	c := coins(prs, [4]int{a, prs[r-1], l, r - 1}, memo) + b - a - 2
	if c < min {
		min = c
	}

	memo[v] = min
	return min
}

func solve(pr *codejam.Problem, d *data) {
	memo := make(map[[4]int]int)

	v := [4]int{0, d.nc + 1, 0, len(d.released)}

	ng := coins(d.released, v, memo)

	pr.Write(fmt.Sprintf("Case #%d: %d\n", d.testIndex, ng))
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
