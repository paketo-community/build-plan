package noapp

import (
	"github.com/buildpack/libbuildpack/layers"
	"github.com/cloudfoundry/libcfbuildpack/build"
)

const Name = "no-app"

type Contributor struct {
	context build.Build
}

func NewContributor(context build.Build) (Contributor, bool, error) {
	_, wantDependency, err := context.Plans.GetShallowMerged(Name)
	if err != nil {
		return Contributor{}, false, err
	}

	if !wantDependency {
		return Contributor{}, false, nil
	}

	return Contributor{context: context}, true, nil
}

func (c Contributor) Contribute() error {
	return c.context.Layers.WriteApplicationMetadata(layers.Metadata{
		Processes: []layers.Process{
			{
				Type:    "web",
				Command: "/bin/bash",
				Direct:  false,
			},
		},
	})
}
