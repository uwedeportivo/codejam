// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/189252/dashboard#s=p0

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/uwedeportivo/codejam"
)

type data struct {
	testIndex int
	val       string
}

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	d := &data{
		testIndex: testCaseIndex,
	}

	d.val = pr.ReadString()

	return d
}

func index(c byte) int {
	switch {
	case c >= 'a' && c <= 'z':
		return int(c - 'a')
	case c >= '0' && c <= '9':
		return int(26 + c - '0')
	default:
		panic("unexpected string character")
	}
}

func solve(pr *codejam.Problem, d *data) {
	var lookup [36]int

	for i := 0; i < 36; i++ {
		lookup[i] = -1
	}

	n := len(d.val)
	if n == 1 {
		pr.Write(fmt.Sprintf("Case #%d: 1\n", d.testIndex))
		return
	}

	countUniq := 0

	for i := 0; i < n; i++ {
		c := d.val[i]
		j := index(c)
		if lookup[j] == -1 {
			if countUniq == 0 {
				lookup[j] = 1
			} else if countUniq == 1 {
				lookup[j] = 0
			} else {
				lookup[j] = countUniq
			}
			countUniq++
		}
	}

	p := 1
	r := 0

	var base int

	if countUniq == 1 {
		base = 2
	} else {
		base = countUniq
	}

	for i := n - 1; i >= 0; i-- {
		c := d.val[i]
		j := index(c)
		if lookup[j] == -1 {
			panic("digit not set")
		}
		r += p * lookup[j]
		p *= base
	}

	pr.Write(fmt.Sprintf("Case #%d: %d\n", d.testIndex, r))
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
