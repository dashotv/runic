package runic

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/dotenv-org/godotenvvault"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := godotenvvault.Load(); err != nil {
		panic(err)
	}
}

func TestRunic_Jackett(t *testing.T) {
	r := &Runic{}
	err := r.Jackett(os.Getenv("JACKETT_URL"), os.Getenv("JACKETT_KEY"))
	assert.NoError(t, err, "runic.New() should not return an error")

	for _, indexer := range r.jackett.indexers {
		fmt.Printf("%s: %s: %s\n", indexer.ID, indexer.Title, indexer.Link)
	}
}

func TestRunic_GeekRead(t *testing.T) {
	r := &Runic{}
	err := r.Add("geek", os.Getenv("NZBGEEK_URL"), os.Getenv("NZBGEEK_KEY"), 0, false)
	assert.NoError(t, err, "runic.Add() should not return an error")

	results, err := r.Read("geek", []int{5000})
	assert.NoError(t, err, "runic.Read() should not return an error")
	assert.NotNil(t, results, "runic.Read() should return a non-nil value")

	for _, result := range results {
		if result.TVTitle != "" {
			fmt.Printf("%s %sx%s %s\n", result.TVTitle, result.Season, result.Episode, result.DownloadURL)
		} else {
			fmt.Printf("%s: %s\n", result.Title, result.DownloadURL)
		}
	}
}

func TestRunic_JackettRead(t *testing.T) {
	r := &Runic{}
	err := r.Jackett(os.Getenv("JACKETT_URL"), os.Getenv("JACKETT_KEY"))
	assert.NoError(t, err, "runic.New() should not return an error")

	results, err := r.Read("nyaasi", []int{5000})
	assert.NoError(t, err, "runic.Read() should not return an error")
	assert.NotNil(t, results, "runic.Read() should return a non-nil value")

	for _, result := range results {
		fmt.Printf("%25.25s: %d %s %s\n", result.Title, result.Size, result.PubDate, strings.Join(result.Category, ","))
	}
}

func TestRunic_NyaaRead(t *testing.T) {
	r := &Runic{}
	err := r.AddTorznab("nyaa", "http://10.0.4.62:9117/api/v2.0/indexers/nyaasi/results/torznab", os.Getenv("JACKETT_KEY"), 0, false)
	assert.NoError(t, err)

	results, err := r.Read("nyaa", []int{5000})
	assert.NoError(t, err, "runic.Read() should not return an error")
	assert.NotNil(t, results, "runic.Read() should return a non-nil value")

	for _, result := range results {
		fmt.Printf("%s: %d: %s\n", result.Title, result.Size, result.DownloadURL)
	}
}
