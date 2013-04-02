// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Parser func(testCaseIndex int, input chan string, output chan interface{})
type Executor func(input chan interface{}, output chan string, done chan bool)

func ParseInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return n
}

func ParseInts(str string, xs []int) {
	vStrs := strings.Split(str, " ")
	if len(vStrs) != len(xs) {
		panic(fmt.Errorf("line: %s did not yield enough numbers", str))
	}

	for i := 0; i < len(xs); i++ {
		v, err := strconv.Atoi(vStrs[i])
		if err != nil {
			panic(err)
		}
		xs[i] = v
	}
}

func ReadLines(filename string, c chan string) {
	inF, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer inF.Close()

	inR := bufio.NewReader(inF)

	for {
		line, err := inR.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)

		c <- line
	}
}

func WriteStrings(filename string, c chan string, done chan bool, doneFlushing chan bool) {
	outF, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer outF.Close()

	writer := bufio.NewWriter(outF)

	for {
		select {
		case line := <-c:
			_, err = writer.WriteString(line)

			if err != nil {
				panic(err)
			}
		case <-done:
			err = writer.Flush()

			if err != nil {
				panic(err)
			}

			doneFlushing <- true
			break
		}
	}
}
