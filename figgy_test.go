package figgy

import (
	"os"
	"testing"
)

func TestFiggy(t *testing.T) {
	os.Setenv("FIGGY_TEST", "hello")
	config := Get(Figgy{"FOO": "bar"}, "FIGGY")

	if len(config) != 2 {
		t.Fatal("Did not get two config values")
	}
}
