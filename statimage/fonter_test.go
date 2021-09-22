package statimage_test

import (
	"os"
	"testing"

	"github.com/rudes/runestats/statapi"
	"github.com/rudes/runestats/statimage"
)

// TestNewRuneStat tests the encoding of an image with a real players stats
func TestNewRuneStat(t *testing.T) {
	_staticDir := "/app/templates/static/"
	os.MkdirAll(_staticDir+"images/os_rs/", os.ModeDir|os.ModePerm)
	player := "peonpower"
	stats := statapi.OldSchoolAPIHandler(player)
	err := statimage.NewRuneStat(player, stats, _staticDir)
	if err != nil {
		t.Errorf("Error creating image : %s", err)
	}
}
