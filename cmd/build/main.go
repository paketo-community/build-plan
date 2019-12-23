package main

import (
	"github.com/cloudfoundry/no-app-cnb/noapp"
	"github.com/cloudfoundry/packit"
)

func main() {
	packit.Build(noapp.Build())
}
