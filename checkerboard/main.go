// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/7214486/dashboard#s=p0

// key insights:
//
// 1.) if one row is checkered, all other rows must be checkered
//     (assume not: to heal two adjacent zeros or ones would destroy other checkered row)
//
// 2.) same thing for columns
//
// 3.) to get minimal swaps in row: find misplace then find first complement misplace and swap
//
// 4.) first do row zero and check all other rows for checkered
//     then do column zero and check all other columns for checkered

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/uwedeportivo/codejam"
)

type matrix [][]byte

func (m matrix) swapRows(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m matrix) swapColumns(i, j int) {
	for k := 0; k < len(m); k++ {
		m[k][i], m[k][j] = m[k][j], m[k][i]
	}
}

func (m matrix) copy() matrix {
	mr := make([][]byte, len(m))

	for i := 0; i < len(m); i++ {
		mr[i] = make([]byte, len(m))

		copy(mr[i], m[i])
	}
	return mr
}

func (m matrix) rowCheckered(i int) bool {
	if len(m) == 0 {
		return true
	}

	l := m[i][0]
	for j := 1; j < len(m); j++ {
		if m[i][j] != 1-l {
			return false
		}
		l = 1 - l
	}
	return true
}

func (m matrix) columnCheckered(i int) bool {
	if len(m) == 0 {
		return true
	}

	l := m[0][i]
	for j := 1; j < len(m); j++ {
		if m[j][i] != 1-l {
			return false
		}
		l = 1 - l
	}
	return true
}

// returns (number of swaps, true) if possible or (0, false) if impossible
func (m matrix) checkerRows(first byte) (int, bool) {
	if len(m) == 0 {
		return 0, true
	}

	ns := 0

	l1 := first
	for i := 0; i < len(m); i++ {
		if m[0][i] != 1-l1 {
			l2 := 1 - l1
			healed := false
			for j := i + 1; j < len(m); j++ {
				if m[0][j] != 1-l2 && m[0][j] == 1-l1 {
					m.swapColumns(i, j)
					ns++
					healed = true
					break
				}
				l2 = 1 - l2
			}
			if !healed {
				return 0, false
			}
		}
		l1 = 1 - l1
	}

	// now check that all other rows are checkered
	for k := 1; k < len(m); k++ {
		if !m.rowCheckered(k) {
			return 0, false
		}
	}
	return ns, true
}

// returns (number of swaps, true) if possible or (0, false) if impossible
func (m matrix) checkerColumns(first byte) (int, bool) {
	if len(m) == 0 {
		return 0, true
	}

	ns := 0

	l1 := first
	for i := 0; i < len(m); i++ {
		if m[i][0] != 1-l1 {
			l2 := 1 - l1
			healed := false
			for j := i + 1; j < len(m); j++ {
				if m[j][0] != 1-l2 && m[j][0] == 1-l1 {
					m.swapRows(i, j)
					ns++
					healed = true
					break
				}
				l2 = 1 - l2
			}
			if !healed {
				return 0, false
			}
		}
		l1 = 1 - l1
	}

	// now check that all other columns are checkered
	for k := 1; k < len(m); k++ {
		if !m.columnCheckered(k) {
			return 0, false
		}
	}
	return ns, true
}

type data struct {
	testIndex int
	m         matrix
}

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	d := &data{
		testIndex: testCaseIndex,
	}

	n := pr.ReadInt()

	d.m = make([][]byte, 2*n)

	for i := 0; i < 2*n; i++ {
		d.m[i] = make([]byte, 2*n)

		str := pr.ReadString()

		for j := 0; j < 2*n; j++ {
			if str[j] == '1' {
				d.m[i][j] = 1
			} else {
				d.m[i][j] = 0
			}
		}
	}
	return d
}

func solveF(m matrix, firstRow, firstColumn byte) (int, bool) {
	mc := m.copy()
	nr, ok := mc.checkerRows(firstRow)
	if !ok {
		return 0, false
	}
	nc, ok := mc.checkerColumns(firstColumn)
	if !ok {
		return 0, false
	}
	return nr + nc, true
}

func solve(pr *codejam.Problem, d *data) {
	// bigger than any number of swaps
	minZero := 4*len(d.m) + 1

	ns := minZero
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			n, ok := solveF(d.m, byte(i), byte(j))
			if ok && n < ns {
				ns = n
			}
		}
	}

	if ns == minZero {
		pr.Write(fmt.Sprintf("Case #%d: IMPOSSIBLE\n", d.testIndex))
	} else {
		pr.Write(fmt.Sprintf("Case #%d: %d\n", d.testIndex, ns))
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
