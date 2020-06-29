package app

import "github.com/blang/semver/v4"

const Version = "v0.0.0"

// GetVersion returns the current version as a semantic version, and panics if it is not a valid format
func GetVersion() semver.Version {
	return semver.MustParse(Version)
}
