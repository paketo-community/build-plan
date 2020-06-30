package buildplan

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/paketo-buildpacks/packit"
)

type BuildPlanParser struct{}

func NewBuildPlanParser() BuildPlanParser {
	return BuildPlanParser{}
}

func (p BuildPlanParser) Parse(path string) ([]packit.BuildPlanRequirement, []packit.BuildPlanRequirement, error) {
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil, nil
		}

		return nil, nil, fmt.Errorf("failed to read plan.toml: %w", err)
	}
	defer file.Close()

	var plan packit.BuildPlan
	_, err = toml.DecodeReader(file, &plan)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode plan.toml: %w", err)
	}

	var orRequirements []packit.BuildPlanRequirement

	for _, buildPlan := range plan.Or {
		orRequirements = append(orRequirements, buildPlan.Requires...)
	}

	return plan.Requires, orRequirements, nil
}
