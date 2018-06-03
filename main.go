package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "github.com/ImpSy/UniversalConverter/converter"
)

// import (
// 	. "github.com/ImpSy/UniversalConverter/converter"
// )

func main() {
	content, err := ioutil.ReadFile("testdata/test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := LoadYaml(string(content))
	res := DumpYaml(data)
	fmt.Fprintf(os.Stdout, res)
}
