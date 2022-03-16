package main

import (
	"github.com/paketo-buildpacks/packit/v2"
	buildplan "github.com/paketo-community/build-plan"
)

func main() {
	planParser := buildplan.NewBuildPlanParser()

	packit.Run(
		buildplan.Detect(planParser),
		buildplan.Build(),
	)
}
