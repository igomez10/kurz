package kurz

import (
	"strings"
	"testing"
)

func TestEncodeURL(t *testing.T) {
	initialURL := "http://google.com"

	db := newLocalDB()
	shortener := newURLShortener(db)
	encoded, err := shortener.EncodeURL(initialURL)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	expectedEncoded := "I0mIVmyaCpz5Us7IKxQ7+cIHrBY="

	if strings.Compare(encoded, expectedEncoded) != 0 {
		baseMessage := "Unexpected value after encoding %s - expected %s got %s"
		t.Errorf(baseMessage, initialURL, expectedEncoded, encoded)
	}
}

func TestDecodeURL(t *testing.T) {

}
