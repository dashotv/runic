package runic

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestNewznab_Read(t *testing.T) {
	files := []string{
		"fixtures/newznab.xml",
	}
	urls := []string{
		"https://api.nzbgeek.info/rss?t=2000&limit=100&r=eISG7JzxXnmWskK632mjY3CHRylfVuiX",
	}

	for _, path := range files {
		r := &NewznabReader{}
		got, err := r.ReadFile(path)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		// pp.Println(got)
	}

	for _, url := range urls {
		r := &NewznabReader{}
		got, err := r.ReadURL(url)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		// pp.Println(got)
	}
}
