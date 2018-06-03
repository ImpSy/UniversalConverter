package converter

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"gopkg.in/yaml.v2"
)

// LoadYaml wip
func LoadYaml(data string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	if err := yaml.Unmarshal([]byte(data), &m); err != nil {
		log.Fatalf("error: %v", err)
	}

	return m
}

// DumpYaml wip
func DumpYaml(m map[interface{}]interface{}) string {
	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return string(d)
}

// LoadToml wip
func LoadToml(data string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	if err := toml.Unmarshal([]byte(data), &m); err != nil {
		log.Fatalf("error: %v", err)
	}

	return m
}

// DumpToml wip
func DumpToml(m map[interface{}]interface{}) string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(m); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

// LoadJSON wip
func LoadJSON(data string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		log.Fatalf("error: %v", err)
	}
	return m
}

// DumpJSON wip
func DumpJSON(m map[interface{}]interface{}) string {
	json, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(json)
}

// LoadHCL wip
func LoadHCL(data string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	if err := hcl.Unmarshal([]byte(data), &m); err != nil {
		log.Fatalf("error: %v", err)
	}
	return m
}

// DumpHCL wip
func DumpHCL(m map[interface{}]interface{}) string {
	ast, err := jsonParser.Parse([]byte(DumpJSON(m)))
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
