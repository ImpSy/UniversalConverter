package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ImpSy/UniversalConverter/converter"
)

var (
	input  = flag.String("i", "", "input")
	output = flag.String("format", "json", "[json, hcl, yaml, toml]")
)

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	c := converter.InitConverter()
	extension := parseExtension(*input)
	c.Load(string(content), extension)
	res := c.Dump(*output)
	fmt.Fprintf(os.Stdout, res)
}

func parseExtension(filename string) string {
	inputSplit := strings.Split(filename, ".")
	if len(inputSplit) > 2 {
		return "no extension found"
	}

	extension := inputSplit[len(inputSplit)-1]
	switch extension {
	case "tf", "tfvars":
		return "hcl"
	case "yml":
		return "yaml"
	default:
		return extension
	}
}
