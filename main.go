package main

import (
  "runtime"
  "github.com/gabehoban/go-cli-template/cmd"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
  cmd.Execute()
}