// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/1836486/dashboard#s=p0
// writeup: doc/safety.pdf

package main

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"os"

	"strconv"

	"github.com/uwedeportivo/codejam"
)

const epsilon = 0.000001

type data struct {
	points    []float64
	testIndex int
	sum       float64
	n         float64
	denom     float64
}

func (d *data) calc(k int) float64 {
	v := d.points[k]

	nominator := 2.0*d.sum - d.n*v

	if nominator < 0 {
		return 0.0
	}

	return nominator / d.denom
}

func (d *data) search(k int) float64 {
	var left, right, alpha float64

	alpha = d.calc(k)
	left = 0.0
	right = 1.0

	for {
		c := d.points[k] + alpha*d.sum

		var total float64

		for i, v := range d.points {
			if i != k && c >= v {
				total += (c - v) / d.sum
			}
		}

		if math.Abs(alpha+total-1.0) < epsilon || left == right {
			break
		} else {
			if alpha+total > 1.0 {
				right = alpha
				alpha = left + (alpha-left)/2.0
			} else {
				left = alpha
				alpha = alpha + (right-alpha)/2.0
			}
		}
	}

	return alpha
}

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	ps := pr.ReadInts(nil)

	d := &data{
		testIndex: testCaseIndex,
	}

	d.points = make([]float64, len(ps)-1)

	for i := 0; i < len(d.points); i++ {
		v := float64(ps[i+1])
		d.points[i] = v
		d.sum += v
	}

	d.n = float64(len(d.points))

	d.denom = d.n * d.sum
	return d
}

func solve(pr *codejam.Problem, d *data) {
	alphas := make([]float64, len(d.points))

	for k, _ := range d.points {
		alphas[k] = d.search(k)
	}

	var buf bytes.Buffer

	first := true
	for _, v := range alphas {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.FormatFloat(v*100.0, 'f', 9, 64))
	}

	pr.Write(fmt.Sprintf("Case #%d: %s\n", d.testIndex, buf.String()))
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
