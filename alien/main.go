// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/90101/dashboard#s=p0

package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/uwedeportivo/codejam"
)

type generator struct {
	pattern  string
	buf      []byte
	tree     [][]byte
	indices  []int
	wordSize int
}

func newGenerator(pattern string, wordSize int) *generator {
	g := new(generator)
	g.pattern = pattern
	g.buf = make([]byte, wordSize)
	g.tree = make([][]byte, wordSize)
	g.wordSize = wordSize

	cursor := 0
	for i := 0; i < wordSize; i++ {
		if pattern[cursor] == '(' {
			cursor++
			for pattern[cursor] != ')' {
				g.tree[i] = append(g.tree[i], pattern[cursor])
				cursor++
			}
			cursor++
		} else {
			g.tree[i] = []byte{pattern[cursor]}
			cursor++
		}
	}

	g.indices = make([]int, wordSize)
	return g
}

func (g *generator) advance() bool {
	i := g.wordSize - 1
	for {
		if g.indices[i] < len(g.tree[i])-1 {
			g.indices[i] = g.indices[i] + 1
			return true
		} else {
			g.indices[i] = 0
			i--
			if i < 0 {
				return false
			}
		}
	}
	// unreachable
	return false
}

func (g *generator) word() string {
	for i := 0; i < g.wordSize; i++ {
		g.buf[i] = g.tree[i][g.indices[i]]
	}

	return string(g.buf)
}

func solve(pr *codejam.Problem, testCaseIndex int, pattern string, words []string) {
	fmt.Printf("test case %d: %s\n", testCaseIndex, pattern)

	var nw int

	g := newGenerator(pattern, len(words[0]))

	for ok := true; ok; ok = g.advance() {
		w := g.word()
		index := sort.SearchStrings(words, w)
		if index < len(words) && words[index] == w {
			nw++
		}
	}

	pr.Write(fmt.Sprintf("Case #%d: %d\n", testCaseIndex, nw))

	fmt.Printf("finished test case %d: %d\n", testCaseIndex, nw)
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

	ldn := pr.ReadInts(nil)
	if len(ldn) != 3 {
		panic(fmt.Errorf("invalid input"))
	}

	dictSize := ldn[1]
	numTestCases := ldn[2]

	if numTestCases < 1 {
		panic(fmt.Errorf("no testcases available"))
	}

	words := make([]string, dictSize)
	for i := 0; i < dictSize; i++ {
		words[i] = pr.ReadString()
	}
	sort.Strings(words)

	for testIndex := 1; testIndex <= numTestCases; testIndex++ {
		solve(pr, testIndex, pr.ReadString(), words)
	}

	pr.Close()
}
