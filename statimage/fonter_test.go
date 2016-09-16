package statimage_test

import (
	"os"
	"testing"

	"github.com/rudes/runestats/statapi"
	"github.com/rudes/runestats/statimage"
)

// TestNewRuneStat tests the encoding of an image with a real players stats
func TestNewRuneStat(t *testing.T) {
	_staticDir := os.Getenv("GOPATH") + "/src/github.com/rudes/runestats/templates/static/"
	os.MkdirAll(_staticDir+"images/os_rs/", os.ModeDir)
	player := "peonpower"
	stats := statapi.OldSchoolAPIHandler(player)
	statimage.NewRuneStat(player, stats, _staticDir)
}
