package noapp

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type EnvironmentParser struct{}

func NewEnvironmentParser() EnvironmentParser {
	return EnvironmentParser{}
}

func (e EnvironmentParser) Parse(workingDir string) ([]string, error) {
	envData, err := ioutil.ReadFile(filepath.Join(workingDir, "env.toml"))
	if err != nil {
		return nil, err
	}

	var env struct {
		Deps []string `toml:"deps"`
	}
	if _, err := toml.Decode(string(envData), &env); err != nil {
		return nil, err
	}
	return env.Deps, nil
}
