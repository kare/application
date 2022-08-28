package application

import (
	"fmt"
)

// Meta contains application version information.
type Meta struct {
	BuildTime string `json:"build_time"`
	Version   string `json:"version"`
	Revision  string `json:"revision"`
}

func require(name, value string) error {
	if value == "" {
		return fmt.Errorf("application: Meta.%s is a required value", name)
	}
	return nil
}

// NewMeta creates a new meta version.
func NewMeta(buildTime, version, revision string) (*Meta, error) {
	if err := require("BuildTime", buildTime); err != nil {
		return nil, err
	}
	if err := require("Version", version); err != nil {
		return nil, err
	}
	if err := require("Revision", revision); err != nil {
		return nil, err
	}
	return &Meta{
		BuildTime: buildTime,
		Version:   version,
		Revision:  revision,
	}, nil
}
