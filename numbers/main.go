// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/32016/dashboard#s=p2
// writeup: doc/threeplussqrtfive.pdf

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/uwedeportivo/codejam"
)

type data struct {
	num       int
	testIndex int
}

var hundredsLookup []int = []int{2, 6, 28, 144, 752, 936, 608, 904, 992, 336, 48, 944, 472, 56, 448, 464, 992, 96,
	608, 264, 152, 856, 528, 744, 352, 136, 408, 904, 792, 136, 648, 344, 472, 456, 848, 264, 192, 96, 808, 464, 552,
	456, 528, 344, 952, 336, 208, 904, 592, 936, 248, 744, 472, 856, 248, 64, 392, 96, 8, 664, 952, 56, 528, 944,
	552, 536, 8, 904, 392, 736, 848, 144, 472, 256, 648, 864, 592, 96, 208, 864, 352, 656, 528, 544, 152, 736, 808,
	904, 192, 536, 448, 544, 472, 656, 48, 664, 792, 96, 408, 64, 752, 256, 528}
var hundredsOffset int = 3
var hundredsPeriod int = 100

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	n := pr.ReadInt()

	d := &data{
		num:       n,
		testIndex: testCaseIndex,
	}

	return d
}

func mod10(v int) int {
	r := v % 10
	if r < 0 {
		r = r + 10
	}
	return r
}

func solve(pr *codejam.Problem, d *data) {
	var unitsDigit, tensDigit, hundredsDigit int

	if d.num == 0 {
		unitsDigit, tensDigit, hundredsDigit = 1, 0, 0
	} else if d.num == 1 {
		unitsDigit, tensDigit, hundredsDigit = 5, 0, 0
	} else if d.num == 2 {
		unitsDigit, tensDigit, hundredsDigit = 7, 2, 0
	} else {
		v := (d.num - hundredsOffset) % hundredsPeriod
		u := hundredsLookup[v+hundredsOffset] - 1

		unitsDigit = mod10(u)

		if u >= 10 {
			u = (u - unitsDigit) / 10

			tensDigit = mod10(u)

			if u >= 10 {
				u = (u - tensDigit) / 10

				hundredsDigit = mod10(u)
			}
		}
	}

	pr.Write(fmt.Sprintf("Case #%d: %d%d%d\n", d.testIndex, hundredsDigit, tensDigit, unitsDigit))
}

func main() {
	help := flag.Bool("help", false, "show this message")
	inFile := flag.String("in", "", "input filename (required)")
	outFile := flag.String("out", "", "output filename (stdout if ommitted)")

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

/*

This code was used to generate the lookup, period and offset

type seqPair struct {
	a   int
	b   int
	mod int
}

func (s *seqPair) advance() {
	c := (6*s.b - 4*s.a) % s.mod

	if c < 0 {
		c = c + s.mod
	}

	s.a, s.b = s.b, c
}

func (s *seqPair) same(ss *seqPair) bool {
	return s.a == ss.a && s.b == ss.b
}

func findPeriod(m int) (int, int) {
	s1 := &seqPair{
		a:   2,
		b:   6,
		mod: m,
	}
	s2 := &seqPair{
		a:   2,
		b:   6,
		mod: m,
	}

	for {
		s1.advance()
		s2.advance()
		s2.advance()

		if s1.same(s2) {
			break
		}
	}

	fa, fb := s1.a, s1.b

	p := 1
	loop := make(map[int]bool)

	loop[1000*s1.a+s1.b] = true

	for {
		s1.advance()

		if s1.a == fa && s1.b == fb {
			break
		} else {
			p++
			loop[1000*s1.a+s1.b] = true
		}
	}

	s3 := &seqPair{
		a:   2,
		b:   6,
		mod: m,
	}

	q := 0

	for {
		if loop[1000*s3.a+s3.b] {
			break
		}
		s3.advance()
		q++
	}

	return p, q
}

func generateLookupTable(mod int) {
	fmt.Printf("periodic for mod %d\n", mod)
	period, offset := findPeriod(mod)
	fmt.Printf("offset=%d, period=%d\n", offset, period)

	fmt.Printf("table=[]int{2, 6, ")

	a := 2
	b := 6
	c := (6*b - 4*a) % mod

	if c < 0 {
		c = c + mod
	}

	fmt.Printf("%d, ", c)

	for k := 1; k < period+offset-2; k++ {
		a = b
		b = c
		c = (6*b - 4*a) % mod

		if c < 0 {
			c = c + mod
		}

		fmt.Printf("%d, ", c)
	}

	fmt.Println("}")
}
*/
