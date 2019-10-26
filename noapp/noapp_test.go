package noapp_test

import (
	"testing"

	"github.com/cloudfoundry/libcfbuildpack/buildpackplan"
	"github.com/cloudfoundry/libcfbuildpack/layers"
	"github.com/cloudfoundry/libcfbuildpack/test"
	"github.com/cloudfoundry/no-app-cnb/noapp"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

func TestUnitNoapp(t *testing.T) {
	spec.Run(t, "Noapp", testNoapp, spec.Report(report.Terminal{}))
}

func testNoapp(t *testing.T, when spec.G, it spec.S) {
	var (
		f *test.BuildFactory
	)

	it.Before(func() {
		RegisterTestingT(t)
		f = test.NewBuildFactory(t)

		f.AddPlan(buildpackplan.Plan{Name: noapp.Name})
	})

	when("noapp.NewContributor", func() {
		it("returns true if no-app is in the buildplan", func() {
			_, willContribute, err := noapp.NewContributor(f.Build)
			Expect(err).NotTo(HaveOccurred())
			Expect(willContribute).To(BeTrue())
		})

		it("returns false if noapp is not in the buildplan", func() {
			f.Build.Plans = buildpackplan.Plans{}

			_, willContribute, err := noapp.NewContributor(f.Build)
			Expect(err).NotTo(HaveOccurred())
			Expect(willContribute).To(BeFalse())
		})
	})

	when("Contribute", func() {
		it("sets the start command when only the runtime is used", func() {
			contributor, _, err := noapp.NewContributor(f.Build)
			Expect(err).NotTo(HaveOccurred())
			Expect(contributor.Contribute()).To(Succeed())
			Expect(f.Build.Layers).To(test.HaveApplicationMetadata(layers.Metadata{
				Processes: []layers.Process{
					{
						Type:    "web",
						Command: "/bin/bash",
						Direct:  false,
					},
				},
			}))
		})

	})
}
