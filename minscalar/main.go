// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package main

// problem https://code.google.com/codejam/contest/32016/dashboard#s=p0

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/uwedeportivo/codejam"
)

type vectorData struct {
	xs        []int
	ys        []int
	testIndex int
}

func (vd *vectorData) scalar() int64 {
	n := len(vd.xs)
	var sum int64

	for i := 0; i < n; i++ {
		sum = sum + int64(vd.xs[i])*int64(vd.ys[n-i-1])
	}
	return sum
}

func parse(pr *codejam.Problem, testCaseIndex int) *vectorData {
	n := pr.ReadInt()

	vd := &vectorData{
		xs:        make([]int, n),
		ys:        make([]int, n),
		testIndex: testCaseIndex,
	}

	pr.ReadInts(vd.xs)
	pr.ReadInts(vd.ys)

	return vd
}

func solve(pr *codejam.Problem, vd *vectorData) {
	sort.Ints(vd.xs)
	sort.Ints(vd.ys)

	minSum := vd.scalar()

	pr.Write(fmt.Sprintf("Case #%d: %d\n", vd.testIndex, minSum))
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
		vd := parse(pr, testIndex)
		solve(pr, vd)
	}

	pr.Close()
}
