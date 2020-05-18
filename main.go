package main

import (
	"github.com/paketo-buildpacks/packit"
)

func main() {
	planParser := NewBuildPlanParser()

	packit.Run(Detect(planParser), Build())
}
