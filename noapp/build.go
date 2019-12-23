package noapp

import (
	"github.com/cloudfoundry/packit"
)

func Build() packit.BuildFunc {
	return func(packit.BuildContext) (packit.BuildResult, error) {
		return packit.BuildResult{
			Processes: []packit.Process{
				{
					Type:    "web",
					Command: "/bin/bash",
					Direct:  false,
				},
			},
		}, nil
	}
}
