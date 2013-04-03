// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/uwedeportivo/codejam/milkshakes"
	"github.com/uwedeportivo/codejam/minscalar"
	"github.com/uwedeportivo/codejam/utils"
)

const (
	versionStr = "1.0"
)

func usage() {
	fmt.Fprintf(os.Stderr, "%s version %s, Copyright (c) 2013 Uwe Hoffmann. All rights reserved.\n", os.Args[0], versionStr)
	fmt.Fprintf(os.Stderr, "\t                 %s -in <input filename> -out <output filename> -problem <problemname>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nFlag defaults:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	help := flag.Bool("help", false, "show this message")
	version := flag.Bool("version", false, "show version")

	inFile := flag.String("in", "", "input filename")
	outFile := flag.String("out", "", "output filename")
	problem := flag.String("problem", "", "which problem to solve")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *version {
		fmt.Fprintf(os.Stderr, "%s version %s, Copyright (c) 2013 Uwe Hoffmann. All rights reserved.\n", os.Args[0], versionStr)
		os.Exit(0)
	}

	if len(*inFile) == 0 || len(*outFile) == 0 || len(*problem) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	parsers := make(map[string]utils.Parser)
	executors := make(map[string]utils.Executor)

	parsers["minscalar"] = minscalar.Parse
	parsers["milkshakes"] = milkshakes.Parse
	executors["minscalar"] = minscalar.Execute
	executors["milkshakes"] = milkshakes.Execute

	input := make(chan string)
	output := make(chan string)
	doneWriting := make(chan bool)
	doneFlushing := make(chan bool)

	go utils.ReadLines(*inFile, input)
	go utils.WriteStrings(*outFile, output, doneWriting, doneFlushing)

	numTestCases := utils.ParseInt(<-input)

	if numTestCases < 1 {
		panic(fmt.Errorf("no testcases available"))
	}

	done := make(chan bool, numTestCases)
	cases := make(chan interface{})

	parser := parsers[*problem]
	executor := executors[*problem]

	go executor(cases, output, done)

	for testIndex := 1; testIndex <= numTestCases; testIndex++ {
		parser(testIndex, input, cases)
	}

	for i := 0; i < numTestCases; i++ {
		<-done
	}

	doneWriting <- true

	<-doneFlushing
}
