package convert

import (
	"bytes"
	"fmt"

	hclprinter "github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"github.com/hashicorp/hcl/v2/hclparse"
)

var parser = hclparse.NewParser()

func JsonToHcl(input []byte) []byte {
	bytes, err := convertJsonToHcl(input)
	if err != nil {
		fmt.Errorf("hclTojson() error. %s", err)
	}
	return bytes
}

func convertJsonToHcl(input []byte) ([]byte, error) {
	ast, err := jsonParser.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("unable to parse JSON: %s", err)
	}
	var buf bytes.Buffer
	if err := hclprinter.Fprint(&buf, ast); err != nil {
		return nil, fmt.Errorf("Unable to print HCL: %s", err)
	}

	return buf.Bytes(), nil
}
