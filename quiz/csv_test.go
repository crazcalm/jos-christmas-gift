package quiz

import (
	"path/filepath"
	"testing"
)

func TestReadCSV(t *testing.T) {
	goodData := filepath.Join("test_data", "good.csv")
	badData := filepath.Join("test_data", "bad.csv")

	var tests = []struct {
		Path        string
		ExpectError bool
	}{
		{"", true},
		{"notAFile", true},
		{badData, true},
		{goodData, false},
	}

	for _, test := range tests {
		_, err := ReadCSV(test.Path)

		if test.ExpectError {
			if err == nil {
				t.Error("Was expecting an error, but err == nil ...")
			}
			return
		}

		if err != nil {
			t.Errorf("Was not expecting an err: %s", err.Error())
		}
	}
}
