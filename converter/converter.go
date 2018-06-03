package converter

import (
	"bytes"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Converter
type Converter struct {
	data map[interface{}]interface{}
}

func InitConverter() *Converter {
	return &Converter{
		data: make(map[interface{}]interface{}),
	}
}

// Load
func (c *Converter) Load(data string, format string) {
	switch format {
	case "hcl":
		c.data = loadHCL(data)
	case "json":
		c.data = loadJSON(data)
	case "toml":
		c.data = loadToml(data)
	case "yaml":
		c.data = loadYaml(data)
	default:
		log.Fatalf("Invalid format: %s", format)
	}
}

// Dump
func (c *Converter) Dump(format string) string {
	var res string
	switch format {
	case "hcl":
		res = dumpHCL(c.data)
	case "json":
		res = dumpJSON(c.data)
	case "toml":
		res = dumpToml(c.data)
	case "yaml":
		res = dumpYaml(c.data)
	default:
		log.Fatalf("Invalid format: %s", format)
	}
	return res
}

func loadYaml(data string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	if err := yaml.Unmarshal([]byte(data), &m); err != nil {
		log.Fatalf("error: %v", err)
	}

	return m
}

func dumpYaml(m map[interface{}]interface{}) string {
	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return string(d)
}

func loadToml(data string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	if err := toml.Unmarshal([]byte(data), &m); err != nil {
		log.Fatalf("error: %v", err)
	}

	return m
}

func dumpToml(m map[interface{}]interface{}) string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(m); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func loadJSON(data string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		log.Fatalf("error: %v", err)
	}
	return m
}

func dumpJSON(m map[interface{}]interface{}) string {
	json, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(append(json, '\n'))
}

func loadHCL(data string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	if err := hcl.Unmarshal([]byte(data), &m); err != nil {
		log.Fatalf("error: %v", err)
	}
	return m
}

func dumpHCL(m map[interface{}]interface{}) string {
	ast, err := jsonParser.Parse([]byte(dumpJSON(m)))
	if err != nil {
		log.Fatalf("unable to parse JSON: %s", err)
	}
	buf := new(bytes.Buffer)
	err = printer.Fprint(buf, ast)
	if err != nil {
		log.Fatalf("unable to print HCL: %s", err)
	}
	return buf.String()
}
