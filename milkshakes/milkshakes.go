// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package milkshakes

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/uwedeportivo/codejam/utils"
)

type customer struct {
	satisfied bool
	malted    []int
	unmalted  []int
}

type data struct {
	flavors             int
	customers           []*customer
	customersByUnmalted [][]*customer
	customersByMalted   [][]*customer
	testIndex           int
}

type batch []int

func (b batch) print() string {
	var buf bytes.Buffer

	first := true
	for _, v := range b {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	return buf.String()
}

func (c *customer) isSatisfied(b batch) bool {
	for _, v := range c.unmalted {
		if b[v] == 0 {
			return true
		}
	}

	for _, v := range c.malted {
		if b[v] == 1 {
			return true
		}
	}

	return false
}

func (c *customer) nextFlavorToMalt(cursor int, b batch) int {
	for i := cursor + 1; i < len(c.malted); i++ {
		v := c.malted[i]
		if b[v] == 0 {
			return i
		}
	}
	return -1
}

func (d *data) findUnsatisfiedCustomers() []*customer {
	var r []*customer

	for _, c := range d.customers {
		if !c.satisfied {
			r = append(r, c)
		}
	}
	return r
}

func (d *data) malted(b batch, v int) {
	for _, c := range d.customersByMalted[v] {
		c.satisfied = true
	}

	for _, c := range d.customersByUnmalted[v] {
		c.satisfied = c.isSatisfied(b)
	}
}

func (d *data) unmalted(v int) {
	for _, c := range d.customersByMalted[v] {
		// optimization: we know this is an undo of a malted set at v
		// because it wasn't satisfied
		c.satisfied = false
	}

	for _, c := range d.customersByUnmalted[v] {
		c.satisfied = true
	}
}

func (d *data) search(b batch) bool {
	uscs := d.findUnsatisfiedCustomers()

	if len(uscs) == 0 {
		return true
	}

	for _, c := range uscs {
		cursor := -1
		for {
			cursor = c.nextFlavorToMalt(cursor, b)

			if cursor == -1 {
				break
			}

			v := c.malted[cursor]

			b[v] = 1
			d.malted(b, v)

			found := d.search(b)

			if found {
				return true
			}

			b[v] = 0
			d.unmalted(v)
		}
	}

	return false
}

func Parse(testCaseIndex int, input chan string, output chan interface{}) {
	n := utils.ParseInt(<-input)
	m := utils.ParseInt(<-input)

	d := &data{
		flavors:             n,
		customers:           make([]*customer, m),
		customersByUnmalted: make([][]*customer, n),
		customersByMalted:   make([][]*customer, n),
		testIndex:           testCaseIndex,
	}

	for i := 0; i < m; i++ {
		nums := utils.ParseInts(<-input, nil)

		c := &customer{
			malted:   make([]int, 0, nums[0]),
			unmalted: make([]int, 0, nums[0]),
		}

		for j := 0; j < nums[0]; j++ {
			mv := nums[2*j+2] == 1
			fv := nums[2*j+1] - 1

			if mv {
				c.malted = append(c.malted, fv)
				d.customersByMalted[fv] = append(d.customersByMalted[fv], c)
			} else {
				c.unmalted = append(c.unmalted, fv)
				c.satisfied = true
				d.customersByUnmalted[fv] = append(d.customersByUnmalted[fv], c)
			}
		}

		d.customers[i] = c
	}

	output <- d
}

func Execute(input chan interface{}, output chan string, done chan bool) {
	for {
		item := <-input
		d := item.(*data)

		b := make(batch, d.flavors)

		for _, c := range d.customers {
			if len(c.malted) == 1 && len(c.unmalted) == 0 {
				v := c.malted[0]
				b[v] = 1
				d.malted(b, v)
			}
		}

		found := d.search(b)

		if found {
			output <- fmt.Sprintf("Case #%d: %s\n", d.testIndex, b.print())
		} else {
			output <- fmt.Sprintf("Case #%d: IMPOSSIBLE\n", d.testIndex)
		}
		done <- true
	}
}
