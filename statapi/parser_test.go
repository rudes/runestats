package statapi

import (
	"reflect"
	"testing"
)

// TestOldSchoolAPIHandler tests the data gathering from the OldSchool RS API
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

	actual := OldSchoolAPIHandler("peonpower")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

// TestnewStatFromAPI tests the processing of data from OSRS API
func TestNewStatFromAPI(t *testing.T) {
	expected := "1"
	actual := newStatFromAPI("-1, 1, 0")
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
	expected = ""
	actual = newStatFromAPI("")
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
