// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/7214486/dashboard#s=p1

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/uwedeportivo/codejam"
)

type data struct {
	testIndex int
	xs        []int
}

type swapChain struct {
	a    int
	b    int
	l    int
	next *swapChain
}

func printChain(sc *swapChain) string {
	var buf bytes.Buffer

	buf.WriteString("{")
	first := true
	for s := sc; s != nil; s = s.next {
		if first {
			first = false
		} else {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%+v", *s))
	}
	buf.WriteString("}")

	return buf.String()
}

func swapWays(sc *swapChain) int {
	l := 0

	for s := sc; s != nil; s = s.next {
		l++
	}

	f := 1
	for r := 1; r <= l; r++ {
		f = f * r
	}
	return f
}

func parse(pr *codejam.Problem, testCaseIndex int) *data {
	d := &data{
		testIndex: testCaseIndex,
	}

	n := uint(pr.ReadInt())
	m := 1 << n

	d.xs = make([]int, m)
	d.xs = pr.ReadInts(d.xs)
	return d
}

func powerSorted(xs []int) bool {
	k := len(xs) >> 1
	return xs[k-1]+1 == xs[k]
}

func swap(as, bs []int) {
	for i := 0; i < len(as); i++ {
		as[i], bs[i] = bs[i], as[i]
	}
}

func powerSort(k int, xs []int) ([]*swapChain, bool) {
	if k > len(xs) {
		return []*swapChain{nil}, true
	}

	var is []int

	for i := 0; i < len(xs); i += k {
		if !powerSorted(xs[i : i+k]) {
			is = append(is, i)
		}
	}

	l := k >> 1

	switch {
	case len(is) == 0:
		return powerSort(k<<1, xs)
	case len(is) == 1:
		j := is[0]
		swap(xs[j:j+l], xs[j+l:j+k])

		var scs []*swapChain
		ok := powerSorted(xs[j : j+k])
		if ok {
			scs, ok = powerSort(k<<1, xs)
		}

		swap(xs[j:j+l], xs[j+l:j+k])

		if ok {
			for i, sc := range scs {
				scs[i] = &swapChain{
					a:    j,
					b:    j + l,
					l:    l,
					next: sc,
				}
			}
			return scs, true
		} else {
			return nil, false
		}
	case len(is) == 2:
		var tscs []*swapChain

		for _, j0 := range []int{is[0], is[0] + l} {
			for _, j1 := range []int{is[1], is[1] + l} {
				swap(xs[j0:j0+l], xs[j1:j1+l])

				if powerSorted(xs[is[0]:is[0]+k]) && powerSorted(xs[is[1]:is[1]+k]) {
					scs, ok := powerSort(k<<1, xs)
					if ok {
						for i, sc := range scs {
							scs[i] = &swapChain{
								a:    j0,
								b:    j1,
								l:    l,
								next: sc,
							}
						}
						tscs = append(tscs, scs...)
					}
				}
				swap(xs[j0:j0+l], xs[j1:j1+l])
			}
		}

		if len(tscs) > 0 {
			return tscs, true
		} else {
			return nil, false
		}
	default:
		return nil, false
	}
}

func solve(pr *codejam.Problem, d *data) {
	tscsc, ok := powerSort(2, d.xs)

	if !ok {
		pr.Write(fmt.Sprintf("Case #%d: 0\n", d.testIndex))
	} else {
		ns := 0
		for _, sc := range tscsc {
			ns += swapWays(sc)
		}
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
