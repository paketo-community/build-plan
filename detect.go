package buildplan

import (
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

//go:generate faux --interface PlanParser --output fakes/plan_parser.go
type PlanParser interface {
	Parse(path string) (requirements []packit.BuildPlanRequirement, orRequirements []packit.BuildPlanRequirement, err error)
}

func Detect(planParser PlanParser) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		requirements, orRequirements, err := planParser.Parse(filepath.Join(context.WorkingDir, "plan.toml"))
		if err != nil {
			return packit.DetectResult{}, err
		}

		if len(requirements) == 0 {
			return packit.DetectResult{}, packit.Fail
		}

		plan := packit.BuildPlan{
			Requires: requirements,
		}

		if len(orRequirements) > 0 {
			plan.Or = []packit.BuildPlan{
				{Requires: orRequirements},
			}
		}

		return packit.DetectResult{Plan: plan}, nil
	}
}
