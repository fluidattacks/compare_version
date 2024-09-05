package main

import (
	"strings"

	hashiVersion "github.com/hashicorp/go-version"
)

type golangVersion struct {
	raw    string
	semVer *hashiVersion.Version
}

func (g golangVersion) compare(o golangVersion) int {
	if g.semVer == nil || o.semVer == nil {
		if g.semVer != nil {
			return 1
		}
		if o.semVer != nil {
			return -1
		}
		return strings.Compare(g.raw, o.raw)
	}
	return g.semVer.Compare(o.semVer)
}

// newGolangVersion creates a new golangVersion struct from the given version string.
// If the version string is "(devel)", it returns nil, nil.
// Otherwise, it parses the version string using hashiVersion.NewSemver and returns
// a golangVersion struct with the raw version string and the parsed semantic version.
// If there is an error parsing the version string, it returns nil and the error.
func newGolangVersion(v string) (*golangVersion, error) {
	if v == "(devel)" {
		return nil, nil
	}

	semver, err := hashiVersion.NewSemver(strings.TrimPrefix(v, "go"))
	if err != nil {
		return nil, err
	}

	return &golangVersion{
		raw:    v,
		semVer: semver,
	}, nil
}
