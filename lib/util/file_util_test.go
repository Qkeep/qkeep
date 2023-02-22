package util

import (
	"testing"
)

func TestListAllFolder(t *testing.T) {
	path := "../../test"
	folders, err := ListAllFolder(path)
	if err != nil {
		t.Error(err)
	}

	if len(folders) != 2 {
		t.Error("folder count not match")
	}

	if folders[0] != "activity" {
		t.Error("folder name not match")
	}
}
