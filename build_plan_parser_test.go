package main_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	main "github.com/ForestEckhardt/build-plan"
	"github.com/paketo-buildpacks/packit"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testEnvironmentParser(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		workingDir string

		planParser main.BuildPlanParser
	)

	it.Before(func() {
		var err error
		workingDir, err = ioutil.TempDir("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		planParser = main.NewBuildPlanParser()
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	context("Parse", func() {
		it.Before(func() {
			Expect(ioutil.WriteFile(filepath.Join(workingDir, "plan.toml"), []byte(`
[[requires]]
  name = "python"

	[requires.metadata]
	  build = true

[[requires]]
  name = "ruby"

	[requires.metadata]
	  launch = true
`), os.ModePerm)).To(Succeed())
		})

		it("returns a list of strings", func() {
			requirements, err := planParser.Parse(filepath.Join(workingDir, "plan.toml"))
			Expect(err).NotTo(HaveOccurred())
			Expect(requirements).To(Equal([]packit.BuildPlanRequirement{
				{
					Name: "python",
					Metadata: map[string]interface{}{
						"build": true,
					},
				},
				{
					Name: "ruby",
					Metadata: map[string]interface{}{
						"launch": true,
					},
				},
			}))
		})

		context("there is no plan.toml", func() {
			it.Before(func() {
				Expect(os.Remove(filepath.Join(workingDir, "plan.toml"))).To(Succeed())
			})

			it("returns an empty list of requirements", func() {
				requirements, err := planParser.Parse(filepath.Join(workingDir, "plan.toml"))
				Expect(err).NotTo(HaveOccurred())
				Expect(requirements).To(BeEmpty())
			})
		})

		context("failure cases", func() {
			context("when the plan.toml cannot be opened", func() {
				it.Before(func() {
					Expect(os.Chmod(filepath.Join(workingDir, "plan.toml"), 0000)).To(Succeed())
				})

				it("returns an error", func() {
					_, err := planParser.Parse(filepath.Join(workingDir, "plan.toml"))
					Expect(err).To(MatchError(ContainSubstring("failed to read plan.toml:")))
					Expect(err).To(MatchError(ContainSubstring("permission denied")))
				})
			})

			context("the plan.toml is malformed", func() {
				it.Before(func() {
					Expect(ioutil.WriteFile(filepath.Join(workingDir, "plan.toml"), []byte("%%%"), os.ModePerm)).To(Succeed())
				})

				it("returns an error", func() {
					_, err := planParser.Parse(filepath.Join(workingDir, "plan.toml"))
					Expect(err).To(MatchError(ContainSubstring("failed to decode plan.toml:")))
					Expect(err).To(MatchError(ContainSubstring("bare keys cannot contain '%'")))
				})
			})
		})
	})
}
