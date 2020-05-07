package noapp

import (
	"path/filepath"

	"github.com/cloudfoundry/packit"
)

//go:generate faux --interface PlanParser --output fakes/plan_parser.go
type PlanParser interface {
	Parse(path string) ([]packit.BuildPlanRequirement, error)
}

type BuildPlanMetadata struct {
	Launch bool `toml:"launch"`
}

func Detect(planParser PlanParser) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		requirements, err := planParser.Parse(filepath.Join(context.WorkingDir, "plan.toml"))
		if err != nil {
			return packit.DetectResult{}, err
		}

		if len(requirements) == 0 {
			return packit.DetectResult{}, packit.Fail
		}

		requirements = append(requirements, packit.BuildPlanRequirement{Name: DependencyName})

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: DependencyName},
				},
				Requires: requirements,
			},
		}, nil
	}
}
