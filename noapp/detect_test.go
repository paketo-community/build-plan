package noapp_test

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/no-app-cnb/noapp"
	"github.com/cloudfoundry/no-app-cnb/noapp/fakes"
	"github.com/cloudfoundry/packit"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		workingDir string

		planParser *fakes.PlanParser
		detect     packit.DetectFunc
	)

	it.Before(func() {
		var err error
		workingDir, err = ioutil.TempDir("", "workingDir")
		Expect(err).NotTo(HaveOccurred())

		planParser = &fakes.PlanParser{}

		detect = noapp.Detect(planParser)
	})

	context("there is a plan.toml in the app dir", func() {
		context("there are requirements in the plan.toml", func() {
			it.Before(func() {
				planParser.ParseCall.Returns.BuildPlanRequirementSlice = []packit.BuildPlanRequirement{
					{
						Name: "python",
						Metadata: map[string]interface{}{
							"launch": true,
						},
					},
					{
						Name: "ruby",
						Metadata: map[string]interface{}{
							"build": true,
						},
					},
				}
			})

			it("passes detection and has those deps in its final buildplan", func() {
				result, err := detect(packit.DetectContext{
					WorkingDir: workingDir,
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Plan).To(Equal(packit.BuildPlan{
					Provides: []packit.BuildPlanProvision{
						{Name: "no-app"},
					},
					Requires: []packit.BuildPlanRequirement{
						{
							Name: "python",
							Metadata: map[string]interface{}{
								"launch": true,
							},
						},
						{
							Name: "ruby",
							Metadata: map[string]interface{}{
								"build": true,
							},
						},
						{
							Name: "no-app",
						},
					},
				}))

				Expect(planParser.ParseCall.Receives.Path).To(Equal(filepath.Join(workingDir, "plan.toml")))
			})
		})

		context("failure cases", func() {
			context("when the plan parsing fails", func() {
				it.Before(func() {
					planParser.ParseCall.Returns.Error = errors.New("some parsing error")
				})

				it("returns an error", func() {
					_, err := detect(packit.DetectContext{
						WorkingDir: workingDir,
					})
					Expect(err).To(MatchError("some parsing error"))
				})
			})

			context("there are no requirements in the plan.toml", func() {
				it("fails detection", func() {
					_, err := detect(packit.DetectContext{
						WorkingDir: workingDir,
					})

					Expect(err).To(MatchError(packit.Fail))
				})
			})
		})
	})
}
