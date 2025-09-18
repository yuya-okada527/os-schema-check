package fileutil

import "testing"

func TestCheckExtension_JSON(t *testing.T) {
    // TODO: implement more tests
	if !CheckExtension("sample.json", ".json") {
		t.Fatalf("expected true for .json extension")
	}
}
