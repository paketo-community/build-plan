package fakes

import (
	"sync"

	"github.com/paketo-buildpacks/packit/v2"
)

type PlanParser struct {
	ParseCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Path string
		}
		Returns struct {
			Requirements   []packit.BuildPlanRequirement
			OrRequirements []packit.BuildPlanRequirement
			Err            error
		}
		Stub func(string) ([]packit.BuildPlanRequirement, []packit.BuildPlanRequirement, error)
	}
}

func (f *PlanParser) Parse(param1 string) ([]packit.BuildPlanRequirement, []packit.BuildPlanRequirement, error) {
	f.ParseCall.Lock()
	defer f.ParseCall.Unlock()
	f.ParseCall.CallCount++
	f.ParseCall.Receives.Path = param1
	if f.ParseCall.Stub != nil {
		return f.ParseCall.Stub(param1)
	}
	return f.ParseCall.Returns.Requirements, f.ParseCall.Returns.OrRequirements, f.ParseCall.Returns.Err
}
