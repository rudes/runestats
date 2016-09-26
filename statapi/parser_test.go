package statapi

import (
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

// TestOldSchoolAPIHandler tests the data gathering from the OldSchool RS API
func TestOldSchoolAPIHandler(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", _oldSchoolAPIURL+"peonpower",
		httpmock.NewStringResponder(200, `498294,798,2963352
		671596,48,89593
		587194,44,55936
		569143,60,273804
		498668,63,378167
		392395,67,582704
		429100,44,56563
		458370,59,267849
		497653,48,86819
		422761,57,203525
		-1,1,-1
		468465,52,123800
		165045,55,166960
		173438,61,316441
		314849,50,101617
		245290,59,252715
		-1,1,-1
		-1,1,-1
		-1,1,-1
		-1,1,-1
		-1,1,-1
		334553,23,6859
		-1,1,-1
		-1,1,-1
		-1,-1
		-1,-1
		-1,-1`))
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
