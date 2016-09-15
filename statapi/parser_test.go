package statapi_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/rudes/runestats/statapi"
)

func TestOldSchoolAPIHandler(t *testing.T) {
	var expected []string
	expected = append(expected, "798")
	expected = append(expected, "48")
	expected = append(expected, "44")
	expected = append(expected, "60")
	expected = append(expected, "63")
	expected = append(expected, "67")
	expected = append(expected, "44")
	expected = append(expected, "59")
	expected = append(expected, "48")
	expected = append(expected, "57")
	expected = append(expected, "1")
	expected = append(expected, "52")
	expected = append(expected, "55")
	expected = append(expected, "61")
	expected = append(expected, "50")
	expected = append(expected, "59")
	expected = append(expected, "1")
	expected = append(expected, "1")
	expected = append(expected, "1")
	expected = append(expected, "1")
	expected = append(expected, "1")
	expected = append(expected, "23")
	expected = append(expected, "1")
	expected = append(expected, "1")
	expected = append(expected, "-1")
	expected = append(expected, "-1")
	expected = append(expected, "-1")

	actual := statapi.OldSchoolAPIHandler("peonpower")

	if !reflect.DeepEqual(expected, actual) {
		log.Fatalf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}
