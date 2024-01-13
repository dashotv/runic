package runic

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestTorznab_Read(t *testing.T) {
	files := []string{
		"fixtures/torznab.xml",
	}
	urls := []string{
		"http://10.0.4.62:9117/api/v2.0/indexers/nyaasi/results/torznab/api?apikey=avyanr9ly9qov3c5c7v35jrtmiaergh7&t=search&cat=&q=",
	}

	for _, path := range files {
		r := &TorznabReader{}
		got, err := r.ReadFile(path)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		// pp.Println(got)
	}

	for _, url := range urls {
		r := &TorznabReader{}
		got, err := r.ReadURL(url)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		// pp.Println(got)
	}
}
