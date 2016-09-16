package statimage_test

import (
	"testing"

	"github.com/rudes/runestats/statapi"
	"github.com/rudes/runestats/statimage"
)

func TestNewRuneStat(t *testing.T) {
	player := "niriviaa"
	stats := statapi.OldSchoolAPIHandler(player)
	statimage.NewRuneStat(player, stats)
}
