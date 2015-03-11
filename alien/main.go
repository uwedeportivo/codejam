// Copyright (c) 2015 Uwe Hoffmann. All rights reserved.

// problem: https://code.google.com/codejam/contest/90101/dashboard#s=p0

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/uwedeportivo/codejam"
)

const offset byte = 'a'

type kids [26]*trie

type trie struct {
	ks kids
}

func (t *trie) insert(val []byte) {
	if len(val) == 0 {
		return
	}

	i := val[0] - offset

	if t.ks[i] == nil {
		t.ks[i] = new(trie)
	}
	t.ks[i].insert(val[1:])
}

func (t *trie) count(pattern [][]byte) int {
	if len(pattern) == 0 {
		return 1
	}

	n := 0

	for _, c := range pattern[0] {
		i := c - offset
		if t.ks[i] != nil {
			n += t.ks[i].count(pattern[1:])
		}
	}
	return n
}

func string2Pattern(pattern string, wordSize int) [][]byte {
	if pattern == "" {
		return nil
	}

	tree := make([][]byte, wordSize)

	cursor := 0
	for i := 0; i < wordSize; i++ {
		if pattern[cursor] == '(' {
			cursor++
			for pattern[cursor] != ')' {
				tree[i] = append(tree[i], pattern[cursor])
				cursor++
			}
			cursor++
		} else {
			tree[i] = []byte{pattern[cursor]}
			cursor++
		}
	}
	return tree
}

func solve(pr *codejam.Problem, testCaseIndex int, patternStr string, root *trie, wordSize int) {
	pattern := string2Pattern(patternStr, wordSize)

	nw := root.count(pattern)

	pr.Write(fmt.Sprintf("Case #%d: %d\n", testCaseIndex, nw))
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

	wordSize := ldn[0]
	dictSize := ldn[1]
	numTestCases := ldn[2]

	if numTestCases < 1 {
		panic(fmt.Errorf("no testcases available"))
	}

	root := new(trie)
	for i := 0; i < dictSize; i++ {
		root.insert([]byte(pr.ReadString()))
	}

	for testIndex := 1; testIndex <= numTestCases; testIndex++ {
		solve(pr, testIndex, pr.ReadString(), root, wordSize)
	}

	pr.Close()
}
