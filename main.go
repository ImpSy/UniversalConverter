package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ImpSy/UniversalConverter/converter"
)

func main() {
	content, err := ioutil.ReadFile("testdata/test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	c := converter.InitConverter()
	c.Load(string(content), "yaml")
	res := c.Dump("hcl")
	fmt.Fprintf(os.Stdout, res)
}
