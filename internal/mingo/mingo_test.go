package mingo

import (
	"embed"
	"fmt"
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata
var testdata embed.FS

func Test_Minify(t *testing.T) {
	dirs, err := testdata.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	for _, dir := range dirs {
		t.Run(dir.Name(), func(t *testing.T) {
			src, err := testdata.ReadFile(fmt.Sprintf("testdata/%s/main.go", dir.Name()))
			if err != nil {
				t.Fatal(err)
			}

			want, err := testdata.ReadFile(fmt.Sprintf("testdata/%s/expected", dir.Name()))
			if err != nil {
				t.Fatal(err)
			}

			got, err := Minify("main.go", src)
			assert.NoError(t, err)
			assert.Equal(t, string(want), got)

			// check syntax
			_, err = format.Source([]byte(got))
			assert.NoError(t, err)
		})
	}
}
