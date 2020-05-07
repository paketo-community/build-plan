package noapp

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/cloudfoundry/packit"
)

type BuildPlanParser struct{}

func NewBuildPlanParser() BuildPlanParser {
	return BuildPlanParser{}
}

func (p BuildPlanParser) Parse(path string) ([]packit.BuildPlanRequirement, error) {
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to read plan.toml: %w", err)
	}
	defer file.Close()

	var plan packit.BuildPlan
	_, err = toml.DecodeReader(file, &plan)
	if err != nil {
		return nil, fmt.Errorf("failed to decode plan.toml: %w", err)
	}

	return plan.Requires, nil
}
