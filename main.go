package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/mjsso/hcljson/convert"
	"github.com/tidwall/pretty"
)

const (
	colorCyan    = "\033[1;36m%s\033[0m"
	colorYellow  = "\033[1;33m%s\033[0m"
	WarningColor = "\033[1;33m[Warning]\033[0m"
	ErrorColor   = "\033[1;31m[Error]\033[0m"
)

func main() {
	js.Module.Get("exports").Set("HclToJson", convert.HclToJson)
	js.Module.Get("exports").Set("JsonToHcl", convert.JsonToHcl)
	// convertHclTest()
}

func convertHclTest() {

	var files []string

	root := "/home/smj/mjParser/test-samples"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		dir, filename := filepath.Split(file)
		if filepath.Ext(file) == ".tf" {
			var fileBytes []byte
			fileBytes, err = ioutil.ReadFile(file)
			converted, err := convert.HclToJson(fileBytes, filename)
			if err != nil {
				fmt.Println(ErrorColor, "Failed to convert file: ", err)
			}

			var indented bytes.Buffer
			if err := json.Indent(&indented, converted, "", "    "); err != nil {
				fmt.Println(ErrorColor, "Failed to indent file: ", err)
			}
			target := os.Stdout
			var outputFileName = strings.Replace(filename, ".tf", ".json", -1)
			var outputFileDir = filepath.Join(dir, "output")
			if _, err := os.Stat(outputFileDir); os.IsNotExist(err) {
				os.Mkdir(outputFileDir, 0755)
			}
			var outputFile = filepath.Join(outputFileDir, outputFileName)
			target, err = os.OpenFile(outputFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModePerm)
			// json 보기위해 넣은 pretty. 구현할 땐 필요없음
			var prettyJson = pretty.Pretty(converted)
			fmt.Fprintf(target, "%s\n", prettyJson)
		}
	}
}
