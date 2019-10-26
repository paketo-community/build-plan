package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/buildpack/libbuildpack/detect"
	"github.com/cloudfoundry/libcfbuildpack/test"
	"github.com/cloudfoundry/no-app-cnb/noapp"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	"github.com/sclevine/spec/report"
)

func TestUnitDetect(t *testing.T) {
	spec.Run(t, "Detect", testDetect, spec.Report(report.Terminal{}))
}

func testDetect(t *testing.T, when spec.G, it spec.S) {
	var factory *test.DetectFactory
	var envFile string

	it.Before(func() {
		RegisterTestingT(t)
		factory = test.NewDetectFactory(t)
		envFile = filepath.Join(factory.Detect.Application.Root, "env.toml")
	})

	when("there is a env.toml in the app dir", func() {
		when("there are deps in the env.toml", func() {
			it("passes detection and has those deps in its final buildplan", func() {
				Expect(ioutil.WriteFile(envFile, []byte(`deps = ["python", "ruby"]`), os.ModePerm)).To(Succeed())
				code, err := runDetect(factory.Detect)
				Expect(err).NotTo(HaveOccurred())
				Expect(code).To(Equal(detect.PassStatusCode))

				Expect(factory.Plans.Plan.Provides).To(Equal([]buildplan.Provided{{Name: "no-app"}}))
				Expect(factory.Plans.Plan.Requires).To(Equal([]buildplan.Required{
					{Name: noapp.Name},
					{Name: "python",
						Version:  "default",
						Metadata: buildplan.Metadata{"launch": true}},
					{Name: "ruby",
						Version:  "default",
						Metadata: buildplan.Metadata{"launch": true}},
				}))
			})
		})

		when("there are no dep in the env.toml", func() {
			it("fails detection", func() {
				Expect(ioutil.WriteFile(envFile, []byte(``), os.ModePerm)).To(Succeed())
				code, err := runDetect(factory.Detect)
				Expect(err).NotTo(HaveOccurred())
				Expect(code).To(Equal(detect.FailStatusCode))
			})
		})
	})

	when("there is no env.toml in the app dir", func() {
		it("fails detection", func() {
			code, err := runDetect(factory.Detect)
			Expect(err).NotTo(HaveOccurred())
			Expect(code).To(Equal(detect.FailStatusCode))
		})
	})
}
