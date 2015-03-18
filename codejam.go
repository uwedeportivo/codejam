// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

/*
Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package codejam

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return n
}

func parseInts(str string, xs []int) []int {
	vStrs := strings.Split(str, " ")
	if xs != nil && len(vStrs) != len(xs) {
		panic(fmt.Errorf("line: %s did not yield enough numbers", str))
	}

	if xs == nil {
		xs = make([]int, len(vStrs))
	}

	for i := 0; i < len(xs); i++ {
		v, err := strconv.Atoi(vStrs[i])
		if err != nil {
			panic(err)
		}
		xs[i] = v
	}
	return xs
}

func parseFloat(str string) float64 {
	r, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}

	return r
}

func parseFloats(str string, xs []float64) []float64 {
	vStrs := strings.Split(str, " ")
	if xs != nil && len(vStrs) != len(xs) {
		panic(fmt.Errorf("line: %s did not yield enough numbers", str))
	}

	if xs == nil {
		xs = make([]float64, len(vStrs))
	}

	for i := 0; i < len(xs); i++ {
		v, err := strconv.ParseFloat(vStrs[i], 64)
		if err != nil {
			panic(err)
		}
		xs[i] = v
	}
	return xs
}

type Problem struct {
	sc   *bufio.Scanner
	fin  *os.File
	fout *os.File
	wr   *bufio.Writer
}

func NewProblem(inFile, outFile string) *Problem {
	pr := new(Problem)

	f, err := os.Open(inFile)
	if err != nil {
		panic(err)
	}
	pr.fin = f

	pr.sc = bufio.NewScanner(pr.fin)

	if len(outFile) > 0 {
		f, err = os.Create(outFile)
		if err != nil {
			panic(err)
		}
		pr.fout = f
	} else {
		pr.fout = os.Stdout
	}
	pr.wr = bufio.NewWriter(pr.fout)

	return pr
}

func (pr *Problem) nextLine() string {
	success := pr.sc.Scan()
	if !success {
		panic(fmt.Errorf("no more line"))
	}
	line := pr.sc.Text()

	return strings.TrimSpace(line)
}

func (pr *Problem) ReadInts(xs []int) []int {
	return parseInts(pr.nextLine(), xs)
}

func (pr *Problem) ReadInt() int {
	return parseInt(pr.nextLine())
}

func (pr *Problem) ReadFloats(xs []float64) []float64 {
	return parseFloats(pr.nextLine(), xs)
}

func (pr *Problem) ReadFloat() float64 {
	return parseFloat(pr.nextLine())
}

func (pr *Problem) ReadString() string {
	return pr.nextLine()
}

func (pr *Problem) Write(str string) {
	_, err := pr.wr.WriteString(str)

	if err != nil {
		panic(err)
	}
}

func (pr *Problem) Close() {
	if err := pr.fin.Close(); err != nil {
		panic(err)
	}
	if err := pr.wr.Flush(); err != nil {
		panic(err)
	}
	if err := pr.fout.Close(); err != nil {
		panic(err)
	}
}
