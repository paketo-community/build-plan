package main

import (
	"github.com/paketo-buildpacks/packit"
)

func Build() packit.BuildFunc {
	return func(packit.BuildContext) (packit.BuildResult, error) {
		return packit.BuildResult{}, nil
	}
}
