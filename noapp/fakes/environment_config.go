package fakes

import "sync"

type EnvironmentConfig struct {
	ParseCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			WorkingDir string
		}
		Returns struct {
			Deps []string
			Err  error
		}
		Stub func(string) ([]string, error)
	}
}

func (f *EnvironmentConfig) Parse(param1 string) ([]string, error) {
	f.ParseCall.Lock()
	defer f.ParseCall.Unlock()
	f.ParseCall.CallCount++
	f.ParseCall.Receives.WorkingDir = param1
	if f.ParseCall.Stub != nil {
		return f.ParseCall.Stub(param1)
	}
	return f.ParseCall.Returns.Deps, f.ParseCall.Returns.Err
}
