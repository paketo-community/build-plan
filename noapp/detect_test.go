package noapp_test

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/cloudfoundry/no-app-cnb/noapp"
	"github.com/cloudfoundry/no-app-cnb/noapp/fakes"
	"github.com/cloudfoundry/packit"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		workingDir string

		environmentConfig *fakes.EnvironmentConfig
		detect            packit.DetectFunc
	)

	it.Before(func() {
		var err error
		workingDir, err = ioutil.TempDir("", "workingDir")
		Expect(err).NotTo(HaveOccurred())

		environmentConfig = &fakes.EnvironmentConfig{}

		detect = noapp.Detect(environmentConfig)
	})

	context("there is a env.toml in the app dir", func() {
		context("there are deps in the env.toml", func() {
			it.Before(func() {
				environmentConfig.ParseCall.Returns.Deps = []string{"python", "ruby"}
			})
			it("passes detection and has those deps in its final buildplan", func() {
				result, err := detect(packit.DetectContext{
					WorkingDir: workingDir,
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Plan).To(Equal(packit.BuildPlan{
					Provides: []packit.BuildPlanProvision{
						{Name: noapp.Name},
					},
					Requires: []packit.BuildPlanRequirement{
						{
							Name: noapp.Name,
						}, {
							Name:    "python",
							Version: "default",
							Metadata: noapp.BuildPlanMetadata{
								LayerFlags: map[string]bool{"launch": true},
							},
						}, {
							Name:    "ruby",
							Version: "default",
							Metadata: noapp.BuildPlanMetadata{
								LayerFlags: map[string]bool{"launch": true},
							},
						},
					},
				}))
			})
		})

		context("failure cases", func() {
			context("when the parse fails", func() {
				it.Before(func() {
					environmentConfig.ParseCall.Returns.Err = errors.New("some parsing error")
				})
				it("returns an error", func() {
					_, err := detect(packit.DetectContext{
						WorkingDir: workingDir,
					})
					Expect(err).To(MatchError("some parsing error"))
				})
			})

			context("there are no dep in the env.toml", func() {
				it("returns an error", func() {
					_, err := detect(packit.DetectContext{
						WorkingDir: workingDir,
					})
					Expect(err).To(MatchError("no dependencies were found in the env.toml"))
				})
			})
		})
	})

	// when("there is no env.toml in the app dir", func() {
	// 	it("fails detection", func() {
	// 		code, err := runDetect(factory.Detect)
	// 		Expect(err).NotTo(HaveOccurred())
	// 		Expect(code).To(Equal(detect.FailStatusCode))
	// 	})
	// })
}
