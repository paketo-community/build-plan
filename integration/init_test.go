package integration_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cloudfoundry/dagger"
	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var (
	buildpack              string
	pythonRuntimeBuildpack string
)

func TestIntegration(t *testing.T) {
	Expect := NewWithT(t).Expect

	root, err := dagger.FindBPRoot()
	Expect(err).ToNot(HaveOccurred())

	buildpack, err = dagger.PackageBuildpack(root)
	Expect(err).NotTo(HaveOccurred())

	pythonRuntimeBuildpack, err = dagger.GetLatestBuildpack("python-runtime-cnb")
	Expect(err).ToNot(HaveOccurred())

	// HACK: we need to fix dagger and the package.sh scripts so that this isn't required
	buildpack = fmt.Sprintf("%s.tgz", buildpack)

	defer func() {
		Expect(dagger.DeleteBuildpack(buildpack)).To(Succeed())
		Expect(dagger.DeleteBuildpack(pythonRuntimeBuildpack)).To(Succeed())
	}()

	SetDefaultEventuallyTimeout(10 * time.Second)

	suite := spec.New("Integration", spec.Report(report.Terminal{}))
	suite("Default", testDefault, spec.Parallel())
	suite.Run(t)
}

func ContainerLogs(id string) func() string {
	docker := occam.NewDocker()

	return func() string {
		logs, _ := docker.Container.Logs.Execute(id)
		return logs.String()
	}
}
