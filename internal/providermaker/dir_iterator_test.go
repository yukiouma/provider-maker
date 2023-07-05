package providermaker

import (
	"testing"
)

func TestFileIterator(t *testing.T) {
	root := "/root/playground/golang/misc-playground/provider-maker/examples/demo"
	expect := 6
	di, err := NewDirIterator(root, DefaultFilters...)
	if err != nil {
		t.Fatal(err)
	}
	dirs := make([]string, 0, 10)
	for {
		p, ok := di.Next()
		if !ok {
			break
		}
		dirs = append(dirs, p)
	}
	if result := len(dirs); result != expect {
		t.Fatalf("expect %d, got %d", expect, result)
	}
}
