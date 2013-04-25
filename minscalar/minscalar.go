// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package minscalar

// problem https://code.google.com/codejam/contest/32016/dashboard#s=p0

import (
	"fmt"
	"sort"

	"github.com/uwedeportivo/codejam/utils"
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

func Parse(testCaseIndex int, input chan string, output chan interface{}) {
	n := utils.ParseInt(<-input)

	vd := &vectorData{
		xs:        make([]int, n),
		ys:        make([]int, n),
		testIndex: testCaseIndex,
	}

	utils.ParseInts(<-input, vd.xs)
	utils.ParseInts(<-input, vd.ys)

	output <- vd
}

func Execute(input chan interface{}, output chan string, done chan bool) {
	for {
		item := <-input
		vd := item.(*vectorData)

		sort.Ints(vd.xs)
		sort.Ints(vd.ys)

		minSum := vd.scalar()

		output <- fmt.Sprintf("Case #%d: %d\n", vd.testIndex, minSum)
		done <- true
	}
}
