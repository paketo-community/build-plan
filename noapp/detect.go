package noapp

import (
	"fmt"

	"github.com/cloudfoundry/packit"
)

//go:generate faux --interface EnvironmentConfig --output fakes/environment_config.go
type EnvironmentConfig interface {
	Parse(workingDir string) (dep []string, err error)
}

type BuildPlanMetadata struct {
	LayerFlags map[string]bool
}

func Detect(environmentConfig EnvironmentConfig) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		buildPlan := packit.BuildPlan{
			Provides: []packit.BuildPlanProvision{
				{Name: Name},
			},
			Requires: []packit.BuildPlanRequirement{
				{Name: Name},
			},
		}

		deps, err := environmentConfig.Parse(context.WorkingDir)
		if err != nil {
			return packit.DetectResult{}, err
		}

		if len(deps) < 1 {
			return packit.DetectResult{}, fmt.Errorf("no dependencies were found in the env.toml")
		}

		for _, dep := range deps {
			buildPlan.Requires = append(buildPlan.Requires, packit.BuildPlanRequirement{
				Name:    dep,
				Version: "default",
				Metadata: BuildPlanMetadata{
					LayerFlags: map[string]bool{"launch": true},
				},
			})
		}

		return packit.DetectResult{
			Plan: buildPlan,
		}, nil
	}
}
