package main

import (
	"github.com/cloudfoundry/no-app-cnb/noapp"
	"github.com/cloudfoundry/packit"
)

func main() {
	planParser := noapp.NewBuildPlanParser()

	packit.Detect(noapp.Detect(planParser))
}
