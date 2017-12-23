package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := Router("", "", "")
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Success case
	res, err := http.Get(ts.URL + "/version")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /version is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	// Method not allowed case
	res, err = http.Post(ts.URL+"/version", "text/plain", nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status code for /version is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusMethodNotAllowed)
	}

	// Not found case
	res, err = http.Get(ts.URL + "/non-exists")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code for /version is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}