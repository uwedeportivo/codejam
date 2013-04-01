// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package minscalar

import (
	"fmt"
	"sort"

	"github.com/uwedeportivo/codejam/utils"
)

type vectorData struct {
	xs []int
	ys []int
}

func (vd *vectorData) scalar() int64 {
	n := len(vd.xs)
	var sum int64

	for i := 0; i < n; i++ {
		sum = sum + int64(vd.xs[i])*int64(vd.ys[n-i-1])
	}
	return sum
}

func Execute(input chan string, output chan string, done chan bool) {

	numTestCases := utils.ParseInt(<-input)

	for testIndex := 1; testIndex <= numTestCases; testIndex++ {
		n := utils.ParseInt(<-input)

		vd := &vectorData{
			xs: make([]int, n),
			ys: make([]int, n),
		}

		utils.ParseInts(<-input, vd.xs)
		utils.ParseInts(<-input, vd.ys)

		sort.Ints(vd.xs)
		sort.Ints(vd.ys)

		minSum := vd.scalar()

		output <- fmt.Sprintf("Case #%d: %d\n", testIndex, minSum)
	}

	done <- true
}
