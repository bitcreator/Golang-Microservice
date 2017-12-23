package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestVersion(t *testing.T)  {
	w := httptest.NewRecorder()

	buildTime := time.Now().Format("20171223_18:04:12")
	commit := "git_hash"
	release := "0.0.1"
	h := version(buildTime, commit, release)
	h(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	info := struct {
		BuildTime string `json:"build_time"`
		Commit string `json:"commit"`
		Release string `json:"release"`
	}{}

	err = json.Unmarshal(greeting, &info)
	if err != nil {
		t.Fatal(err)
	}

	if info.Release != release {
		t.Errorf("Release version is wrong. Have: %s, want: %s.", info.Release, release)
	}

	if info.BuildTime != buildTime {
		t.Errorf("Build time is wrong. Have: %s, want: %s.", info.BuildTime, buildTime)
	}

	if info.Commit != commit {
		t.Errorf("Commit is wrong. Have: %s, want: %s.", info.Commit, commit)
	}
}
