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
