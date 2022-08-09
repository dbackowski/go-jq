package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func findValueAtPath(a map[string]any, path []string) (any, error) {
	var v any = a

	for _, part := range path {
		if arr, ok := v.([]any); ok {
			n, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}

			v = arr[n]
			continue
		}

		m, ok := v.(map[string]any)
		if !ok {
			return nil, nil
		}

		v, ok = m[part]
		if !ok {
			return nil, nil
		}
	}

	return v, nil
}

func parseJsonAndFind(path []string) {
	dec := json.NewDecoder(os.Stdin)
	var input map[string]any
	enc := json.NewEncoder(os.Stdout)

	for {
		err := dec.Decode(&input)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		output, err := findValueAtPath(input, path)
		if err != nil {
			log.Fatal(err)
		}

		err = enc.Encode(output)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func main() {
	if len(os.Args) == 1 {
		fmt.Println("No filter provided")
		return
	}

	path := strings.Split(os.Args[1], ".")
	if path[0] == "" {
		path = path[1:]
	}

	parseJsonAndFind(path)
}
