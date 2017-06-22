package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
}

func printUsage() {
	fmt.Println(
		` Usage: githuber [options] [method]

 Githuber is cli tool for working with github api v3. For documentation about
 github api see https://developer.github.com/v3.

 Options:
  -v (version) shows information about current version
  -h (help) shows this message

 For documentation of specific method use githuber -h method
`)
}