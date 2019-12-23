package noapp_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/no-app-cnb/noapp"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testEnvironmentParser(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		envFile    string
		workingDir string

		envParser noapp.EnvironmentParser
	)

	it.Before(func() {
		var err error
		workingDir, err = ioutil.TempDir("", "workingDir")
		Expect(err).NotTo(HaveOccurred())

		envFile = filepath.Join(workingDir, "env.toml")
		envParser = noapp.NewEnvironmentParser()

	})

	context("Parse", func() {
		context("there is a env.toml in the app dir", func() {
			context("there are deps in the env.toml", func() {
				it.Before(func() {
					Expect(ioutil.WriteFile(envFile, []byte(`deps = ["python", "ruby"]`), os.ModePerm)).To(Succeed())
				})
				it("returns a list of strings", func() {
					deps, err := envParser.Parse(workingDir)
					Expect(err).NotTo(HaveOccurred())
					Expect(deps).To(ConsistOf([]string{"python", "ruby"}))
				})
			})
		})

		context("failure cases", func() {
			context("there is no env.toml", func() {
				it("returns an error", func() {
					_, err := envParser.Parse(workingDir)
					Expect(err).To(MatchError(ContainSubstring("no such file or directory")))
				})
			})

			context("the is no env.toml is malformed", func() {
				it.Before(func() {
					Expect(ioutil.WriteFile(envFile, []byte(`deps = 2`), os.ModePerm)).To(Succeed())
				})
				it("returns an error", func() {
					_, err := envParser.Parse(workingDir)
					Expect(err).To(MatchError(ContainSubstring("cannot load TOML value")))
				})
			})
		})
	})
}
