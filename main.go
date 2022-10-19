package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findValueAtPath(a map[string]any, path []string) (any, error) {
	var v any = a

	for _, part := range path {
		if arr, ok := v.([]any); ok {
			array_index, err := strconv.Atoi(part)
			if err != nil {
				return nil, errors.New("invalid path")
			}

			if array_index > len(arr)-1 {
				return nil, errors.New("path not found")
			}

			v = arr[array_index]
			continue
		}

		m, ok := v.(map[string]any)
		if !ok {
			return nil, nil
		}

		v, ok = m[part]
		if !ok {
			return nil, errors.New("path not found")
		}
	}

	return v, nil
}

func parseJsonAndFind(in io.Reader, path []string) {
	dec := json.NewDecoder(in)
	var input map[string]any

	for {
		err := dec.Decode(&input)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("invalid json")
			return
		}

		output, err := findValueAtPath(input, path)
		if err != nil {
			fmt.Println(err)
			return
		}

		bytes, err := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(bytes))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func argsToPath(args []string) []string {
	if len(args) == 1 {
		return make([]string, 0)
	}

	ereg := regexp.MustCompile(`\[([0-9]+)\]`)
	arg := ereg.ReplaceAllString(args[1], ".$1")
	arg = strings.TrimLeft(arg, ".")

	return strings.Split(arg, ".")
}

func main() {
	path := argsToPath(os.Args)
	parseJsonAndFind(os.Stdin, path)
}
