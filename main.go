package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	name     = "cli-template"
	version  = "0.1.0"
	revision = "HEAD"
)

var printVersion = flag.Bool("version", false, "print version")

// dir file 問わず存在確認に使える
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func fatal(format string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, format, err)
	} else {
		fmt.Fprint(os.Stderr, format)
	}
	os.Exit(1)
}

func usage() {
	fmt.Println(`
nipo <command> [arguments]
	`)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Printf("%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return
	}

	if flag.NArg() == 0 {
		usage()
	}

	switch flag.Arg(0) {
	default:
		usage()
	}
}
