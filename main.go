package main

import (
  "github.com/gabehoban/go-cli-template/cmd"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
  cmd.Execute()
}