package main

import (
	"flag"
	"fmt"
	"os"

	python_version "github.com/aquasecurity/go-pep440-version"
	semver "github.com/hashicorp/go-version"
	apk_version "github.com/knqyf263/go-apk-version"
	debian_version "github.com/knqyf263/go-deb-version"
	maven_version "github.com/masahiro331/go-mvn-version"
	php_version "github.com/mcuadros/go-version"
)

type restrictedLanguage string

var allowedLanguages = []string{"rpm", "java", "golang", "ruby", "gem", "php", "composer", "apk", "deb", "debian", "python", "semver", "maven"}

func (r *restrictedLanguage) Set(value string) error {
	for _, v := range allowedLanguages {
		if value == v {
			*r = restrictedLanguage(value)
			return nil
		}
	}
	return fmt.Errorf("valor no permitido: %s. Valores permitidos son: %v", value, allowedLanguages)
}
func (r *restrictedLanguage) String() string {
	return string(*r)
}

func main() {
	var language restrictedLanguage

	flag.Var(&language, "lang", "Programming language")
	version1 := flag.String("v1", "0", "Version 1")
	version2 := flag.String("v2", "0", "Version 2")
	flag.Parse()

	var result int
	var err error

	switch language {
	case "rpm":
		result, err = compareRpmVersions(*version1, *version2)
	case "java", "maven":
		result, err = compareJavaVersions(*version1, *version2)
	case "golang":
		result, err = compareGolangVersions(*version1, *version2)
	case "ruby", "gem":
		result, err = compareRubyGemVersions(*version1, *version2)
	case "php", "composer":
		result, err = comparePhpVersion(*version1, *version2)
	case "apk":
		result, err = compareApkVersions(*version1, *version2)
	case "deb", "debian":
		result, err = compareDebianVersions(*version1, *version2)
	case "python":
		result, err = comparePythonVersions(*version1, *version2)
	case "semver":
		result, err = compareSemanticVersions(*version1, *version2)
	default:
		fmt.Printf("Unsupported language: %s\n", language)
		return
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error comparing versions: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", result)
}

func compareRubyGemVersions(v1, v2 string) (int, error) {
	ver1, err := newGemfileVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid RubyGem version %s: %v", v1, err)
	}
	ver2, err := newGemfileVersion(v2)

	if err != nil {
		return 0, fmt.Errorf("invalid RubyGem version %s: %v", v2, err)
	}

	return ver1.verObj.Compare(ver2.verObj), nil
}

func compareRpmVersions(v1, v2 string) (int, error) {
	ver1, err := newRpmVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid RPM version %s: %v", v1, err)
	}
	ver2, err := newRpmVersion(v2)
	if err != nil {
		return 0, fmt.Errorf("invalid RPM version %s: %v", v2, err)
	}
	return ver1.compare(ver2), nil
}

func compareJavaVersions(v1, v2 string) (int, error) {
	ver1, err := maven_version.NewVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid Java version %s: %v", v1, err)
	}
	ver2, err := maven_version.NewVersion(v2)
	if err != nil {
		return 0, fmt.Errorf("invalid Java version %s: %v", v2, err)
	}
	return ver1.Compare(ver2), nil
}

func compareGolangVersions(v1, v2 string) (int, error) {
	ver1, err := newGolangVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid Golang version %s: %v", v1, err)
	}
	ver2, err := newGolangVersion(v2)
	if err != nil {
		return 0, fmt.Errorf("invalid Golang version %s: %v", v2, err)
	}
	return ver1.compare(*ver2), nil
}

func comparePhpVersion(v1, v2 string) (int, error) {
	return php_version.CompareSimple(v1, v2), nil
}

func compareApkVersions(v1, v2 string) (int, error) {
	ver1, err := apk_version.NewVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid Alpine version %s: %v", v1, err)
	}

	ver2, err := apk_version.NewVersion(v2)

	if err != nil {
		return 0, fmt.Errorf("invalid Alpine version %s: %v", v2, err)
	}

	return ver1.Compare(ver2), nil
}

func compareDebianVersions(v1, v2 string) (int, error) {
	ver1, err := debian_version.NewVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid Debian version %s: %v", v1, err)
	}

	ver2, err := debian_version.NewVersion(v2)

	if err != nil {
		return 0, fmt.Errorf("invalid Debian version %s: %v", v1, err)
	}

	return ver1.Compare(ver2), nil
}

func comparePythonVersions(v1, v2 string) (int, error) {
	ver1, err := python_version.Parse(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid Python version %s: %v", v1, err)
	}

	ver2, err := python_version.Parse(v2)

	if err != nil {
		return 0, fmt.Errorf("invalid Python composer version %s: %v", v1, err)
	}

	return ver1.Compare(ver2), nil
}
func compareSemanticVersions(v1, v2 string) (int, error) {
	ver1, err := semver.NewVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid Semantic version %s: %v", v1, err)
	}

	ver2, err := semver.NewVersion(v2)

	if err != nil {
		return 0, fmt.Errorf("invalid Semantic version %s: %v", v1, err)
	}

	return ver1.Compare(ver2), nil
}
