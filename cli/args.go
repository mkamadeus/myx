package main

import "github.com/alexflint/go-arg"

type Args struct {
	Path string `arg:"required"`
}

func parseArgs() error {
	args := &Args{}
	return arg.Parse(args)
}
