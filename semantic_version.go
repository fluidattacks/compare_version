package main

import (
	"fmt"
	"strings"

	hashiVer "github.com/hashicorp/go-version"
)

type semanticVersion struct {
	verObj *hashiVer.Version
}

var normalizer = strings.NewReplacer(".alpha", "-alpha", ".beta", "-beta", ".rc", "-rc")

func newSemanticVersion(raw string) (*semanticVersion, error) {
	verObj, err := hashiVer.NewVersion(normalizer.Replace(raw))
	if err != nil {
		return nil, fmt.Errorf("unable to create semver obj: %w", err)
	}
	return &semanticVersion{
		verObj: verObj,
	}, nil
}
