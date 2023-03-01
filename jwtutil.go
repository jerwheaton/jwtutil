package main

import (
	"github.com/jerwheaton/jwtutil/cmd"
)

func main() {
	c := cmd.RunCommand()
	if err := c.Execute(); err != nil {
		panic(err)
	}
}
