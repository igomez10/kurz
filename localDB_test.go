package kurz

import (
	"strings"
	"testing"
)

func TestLocalDBGet(t *testing.T) {
	db := map[string]string{
		"ABC": "https://google.com",
		"ELK": "https://elastic.com",
	}
	loc := localDB(db)
	google := loc.Get([]byte("ABC"))
	if strings.Compare(string(google), loc["ABC"]) != 0 {
		t.Errorf("Unexpected entry, ABC should poing to http://google.com")
	}
}

func TestLocalDBPut(t *testing.T) {
	db := map[string]string{}
	loc := localDB(db)

	testV := []string{"ABC", "http://google.com"}

	loc.Put([]byte(testV[0]), []byte(testV[1]))

	if strings.Compare(db[testV[0]], testV[1]) != 0 {
		baseMessage := "unexpected value after Put, expected %+v : %+v - but got  %+v \n\t %+v"
		t.Errorf(baseMessage, testV[0], testV[1], db[testV[0]], db)

	}

}
