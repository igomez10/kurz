package kurz

import (
	"strings"
	"testing"
)

func TestGetURLParam(t *testing.T) {
	urlToEncode := "google.com"
	submittedURL := "/encode/Z29vZ2xlLmNvbQo="
	basePath := "/encode"

	extractedURL, err := getURLParam(basePath, submittedURL)
	if err != nil {
		t.Log(err)
		t.Errorf("Unexpected error extracting url from path %s", err)
	}

	if strings.Compare(extractedURL, urlToEncode) != 0 {
		t.Errorf("Expected '%+v' as given url to encode - got '%+v'", urlToEncode, extractedURL)
	}
}
