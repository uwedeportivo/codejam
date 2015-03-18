// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/189252/dashboard#s=p1

package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/uwedeportivo/codejam"
)

type vector [6]float64

type data struct {
	testIndex int
	iPos      []vector
}

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	d := &data{
		testIndex: testCaseIndex,
	}

	n := pr.ReadInt()
	d.iPos = make([]vector, n)

	for i := 0; i < n; i++ {
		pr.ReadFloats(d.iPos[i][:])
	}

	return d
}

func solve(pr *codejam.Problem, d *data) {
	n := len(d.iPos)
	nf := float64(n)

	var sums vector

	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			sums[j] += d.iPos[i][j] / nf
		}
	}

	alphaSq := sums[0]*sums[0] + sums[1]*sums[1] + sums[2]*sums[2]
	betaSq := sums[3]*sums[3] + sums[4]*sums[4] + sums[5]*sums[5]
	alphaBeta := sums[0]*sums[3] + sums[1]*sums[4] + sums[2]*sums[5]

	tmin := -alphaBeta / betaSq

	var dmin float64

	if tmin < 0.0 || betaSq < 0.00000001 {
		tmin = 0.0
		dmin = math.Sqrt(alphaSq)
	} else {
		x := sums[0] + sums[3]*tmin
		y := sums[1] + sums[4]*tmin
		z := sums[2] + sums[5]*tmin

		dmin = math.Sqrt(x*x + y*y + z*z)
	}

	pr.Write(fmt.Sprintf("Case #%d: %.8f %.8f\n", d.testIndex, dmin, tmin))
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
