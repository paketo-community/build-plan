package main

import (
	"github.com/cloudfoundry/no-app-cnb/noapp"
	"github.com/cloudfoundry/packit"
)

func main() {
	environmentParser := noapp.NewEnvironmentParser()

	packit.Detect(noapp.Detect(environmentParser))
}
