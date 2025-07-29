package main

import (
	"github.com/Uhm-why/gusher/cmd"
)

type file struct {
	name string
	size uint64
}

func main() {
	cmd.Execute()
}
