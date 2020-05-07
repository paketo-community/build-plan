package fakes

import (
	"sync"

	"github.com/cloudfoundry/packit"
)

type PlanParser struct {
	ParseCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Path string
		}
		Returns struct {
			BuildPlanRequirementSlice []packit.BuildPlanRequirement
			Error                     error
		}
		Stub func(string) ([]packit.BuildPlanRequirement, error)
	}
}

func (f *PlanParser) Parse(param1 string) ([]packit.BuildPlanRequirement, error) {
	f.ParseCall.Lock()
	defer f.ParseCall.Unlock()
	f.ParseCall.CallCount++
	f.ParseCall.Receives.Path = param1
	if f.ParseCall.Stub != nil {
		return f.ParseCall.Stub(param1)
	}
	return f.ParseCall.Returns.BuildPlanRequirementSlice, f.ParseCall.Returns.Error
}
