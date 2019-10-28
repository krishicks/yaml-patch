package yamlpatch_test

import (
	"bytes"
	"io"
	"log"
	"testing"

	yamlpatch "github.com/krishicks/yaml-patch"
	yaml "gopkg.in/yaml.v2"
)

func TestMutidoc(t *testing.T) {
	doc := []byte(`---
foo: bar
baz:
  quux: grault
---
foo: baz
baz:
  quux: what
`)

	ops := []byte(`---
- op: add
  path: /baz/waldo
  value: fred
`)

	patch, err := yamlpatch.DecodePatch(ops)
	if err != nil {
		log.Fatalf("decoding patch failed: %s", err)
	}

	bs, err := patch.Apply(doc)
	if err != nil {
		log.Fatalf("applying patch failed: %s", err)
	}

	docDecoder := yaml.NewDecoder(bytes.NewReader(bs))

	var iface map[string]interface{}
	count := 0
	for {
		err := docDecoder.Decode(&iface)
		if err == io.EOF || iface == nil {
			break
		}
		count++
		if iface["baz"].(map[interface{}]interface{})["waldo"] != "fred" {
			t.Error(`"fred" value was not added`)
		}
	}

	if count != 2 {
		t.Error("result should be two documents")
	}
}
