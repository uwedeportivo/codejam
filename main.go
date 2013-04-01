// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"os"

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

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *version {
		fmt.Fprintf(os.Stderr, "%s version %s, Copyright (c) 2013 Uwe Hoffmann. All rights reserved.\n", os.Args[0], versionStr)
		os.Exit(0)
	}

	if len(*inFile) == 0 || len(*outFile) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	input := make(chan string)
	output := make(chan string)
	done := make(chan bool)

	go utils.ReadLines(*inFile, input)
	go utils.WriteStrings(*outFile, output, done)

	minscalar.Execute(input, output, done)
}
