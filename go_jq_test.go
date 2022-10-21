package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestParseJsonAndFind(t *testing.T) {
	args := []string{"./go-jq", ".repo"}
	path := argsToPath(args)

	in := strings.NewReader("{\"id\":\"123\",\"type\":\"event\",\"repo\":{\"id\":\"2222\",\"type\":\"private\"},\"events\":[{\"id\":\"1\"},{\"id\":\"2\"}],\"test\":[[1,2,3],[4,5,6]]}")

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parseJsonAndFind(in, path)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	got := string(out)
	want := "{\n  \"id\": \"2222\",\n  \"type\": \"private\"\n}\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
