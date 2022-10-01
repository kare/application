//go:build !integration

package build_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"kkn.fi/application/build"
)

func TestBuildHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/version", nil)
	rr := httptest.NewRecorder()
	expectedVersion := "v0.1.0"
	expectedBuildTime := "1.1.2020"
	expectedRevision := "2545eac8be2ec2fa73ad03749cc85646acd78bfe"
	expectedGoVersion := "go1.17.2"
	handler := build.Handler(
		expectedVersion,
		expectedBuildTime,
		expectedRevision,
		expectedGoVersion,
	)
	handler.ServeHTTP(rr, req)
	res := rr.Result()

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := &struct {
		Version   string
		BuildTime string
		Revision  string
		GoVersion string
	}{
		Version:   expectedVersion,
		BuildTime: expectedBuildTime,
		Revision:  expectedRevision,
		GoVersion: expectedGoVersion,
	}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		t.Errorf("error decoding json response: %v", err)
	}
	if response.Version != expectedVersion {
		t.Errorf("handler returned unexpected version: got %v want %v", response.Version, expectedVersion)
	}
	if response.BuildTime != expectedBuildTime {
		t.Errorf("handler returned unexpected build time: got %v want %v", response.BuildTime, expectedBuildTime)
	}
	if response.Revision != expectedRevision {
		t.Errorf("handler returned unexpected version: got %v want %v", response.Revision, expectedRevision)
	}
	if response.GoVersion != expectedGoVersion {
		t.Errorf("handler returned unexpected Go version: got %v want %v", response.GoVersion, expectedGoVersion)
	}
}
