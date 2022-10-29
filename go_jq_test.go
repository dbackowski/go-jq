package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func captureParseJsonAndFindConsoleOutput(input io.Reader, path []string) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parseJsonAndFind(input, path)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}

func TestParseJsonAndFind(t *testing.T) {
	t.Run("for valid json input and path set as '.repo' it returns propers data", func(t *testing.T) {
		path := argsToPath([]string{"./go-jq", ".repo"})
		input := strings.NewReader("{\"id\":\"123\",\"type\":\"event\",\"repo\":{\"id\":\"2222\",\"type\":\"private\"},\"events\":[{\"id\":\"1\"},{\"id\":\"2\"}],\"test\":[[1,2,3],[4,5,6]]}")

		got := captureParseJsonAndFindConsoleOutput(input, path)
		want := "{\n  \"id\": \"2222\",\n  \"type\": \"private\"\n}\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("for valid json input and path set as '.test[0]' it returns propers data", func(t *testing.T) {
		path := argsToPath([]string{"./go-jq", ".test[0]"})
		input := strings.NewReader("{\"id\":\"123\",\"type\":\"event\",\"repo\":{\"id\":\"2222\",\"type\":\"private\"},\"events\":[{\"id\":\"1\"},{\"id\":\"2\"}],\"test\":[[1,2,3],[4,5,6]]}")

		got := captureParseJsonAndFindConsoleOutput(input, path)
		want := "[\n  1,\n  2,\n  3\n]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestArgsToPath(t *testing.T) {
	args := []string{"./go-jq", ".repo"}
	got := argsToPath(args)

	want := []string{"repo"}

	if !slices.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}

	args = []string{"/go-jq", ".repo.type"}
	got = argsToPath(args)
	want = []string{"repo", "type"}

	if !slices.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}

	args = []string{"/go-jq", ".test[0]"}
	got = argsToPath(args)
	want = []string{"test", "0"}

	if !slices.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
