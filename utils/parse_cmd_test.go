package utils_test

import (
	"reflect"
	"testing"

	"github.com/sethigeet/ssh-ell/utils"
)

func TestParseCmd(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"without quotes":                               {input: "parsing works", want: []string{"parsing", "works"}},
		"with one double quote":                        {input: "parsing \"gibberish works", want: []string{"parsing", "gibberish works"}},
		"with two double quotes":                       {input: "parsing \"quoted words\" also works", want: []string{"parsing", "quoted words", "also", "works"}},
		"with two double quotes and word between them": {input: "parsing \"quoted ... word in between ... words\" also works", want: []string{"parsing", "quoted ... word in between ... words", "also", "works"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := utils.ParseCmd(test.input)
			if !reflect.DeepEqual(test.want, got) {
				for i, g := range got {
					t.Logf("%d: %s", i, g)
				}
				t.Fatalf("expected: %v, got: %v", test.want, got)
			}
		})
	}
}
