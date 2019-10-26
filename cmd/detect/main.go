package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/libcfbuildpack/detect"
	"github.com/cloudfoundry/libcfbuildpack/helper"
	"github.com/cloudfoundry/no-app-cnb/noapp"
)

type EnvFile struct {
	Deps []string `toml:"deps"`
}

func main() {
	context, err := detect.DefaultDetect()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create a default detection context: %s", err)
		os.Exit(100)
	}

	code, err := runDetect(context)
	if err != nil {
		context.Logger.Info(err.Error())
	}

	os.Exit(code)
}

func runDetect(context detect.Detect) (int, error) {
	envFile := filepath.Join(context.Detect.Application.Root, "env.toml")
	exist, err := helper.FileExists(envFile)
	if err != nil {
		return detect.FailStatusCode, err
	}

	if !exist {
		return detect.FailStatusCode, nil
	}

	plan := buildplan.Plan{
		Provides: []buildplan.Provided{{Name: noapp.Name}},
		Requires: []buildplan.Required{{Name: noapp.Name}},
	}

	envData, err := ioutil.ReadFile(envFile)
	env := EnvFile{}
	if err != nil {
		return detect.FailStatusCode, err
	}

	if _, err := toml.Decode(string(envData), &env); err != nil {
		return detect.FailStatusCode, err
	}

	if len(env.Deps) < 1 {
		return detect.FailStatusCode, nil
	}

	for _, dep := range env.Deps {
		plan.Requires = append(plan.Requires, buildplan.Required{
			Name:     dep,
			Version:  "default",
			Metadata: buildplan.Metadata{"launch": true},
		})
	}

	return context.Pass(plan)
}
