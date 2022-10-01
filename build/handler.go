package build

import (
	"encoding/json"
	"net/http"
	"runtime"
)

// NewHandler creates a Version hander with Go version set to current runtime.
func NewHandler(version, buildTime, revision string) http.Handler {
	goVersion := runtime.Version()
	return Handler(version, buildTime, revision, goVersion)
}

// Handler handler returns the version, build time, and Go version as a JSON message.
func Handler(version, buildTime, revision, goVersion string) http.Handler {
	type response struct {
		Version   string `json:"version"`
		BuildTime string `json:"build_time"`
		Revision  string `json:"revision"`
		GoVersion string `json:"go_version"`
	}
	res := &response{
		Version:   version,
		BuildTime: buildTime,
		Revision:  revision,
		GoVersion: goVersion,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
