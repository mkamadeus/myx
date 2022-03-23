package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/spec"
)

func main() {
	path := "./examples/spec.yaml"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	s, err := spec.Parse(b)
	if err != nil {
		panic(err)
	}

	err = generator.RenderSpec(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
