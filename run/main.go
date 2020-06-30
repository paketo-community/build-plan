package main

import (
	buildplan "github.com/ForestEckhardt/build-plan"
	"github.com/paketo-buildpacks/packit"
)

func main() {
	planParser := buildplan.NewBuildPlanParser()

	packit.Run(
		buildplan.Detect(planParser),
		buildplan.Build(),
	)
}
