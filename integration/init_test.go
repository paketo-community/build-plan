package integration_test

import (
	"path/filepath"
	"testing"
	"time"

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

	root, err := filepath.Abs("../")
	Expect(err).ToNot(HaveOccurred())

	buildpackStore := occam.NewBuildpackStore()

	buildpack, err = buildpackStore.Get.
		WithVersion("1.2.3").
		Execute(root)
	Expect(err).ToNot(HaveOccurred())

	pythonRuntimeBuildpack, err = buildpackStore.Get.
		Execute("github.com/paketo-community/cpython")
	Expect(err).ToNot(HaveOccurred())

	SetDefaultEventuallyTimeout(10 * time.Second)

	suite := spec.New("Integration", spec.Report(report.Terminal{}))
	suite("Default", testDefault, spec.Parallel())
	suite("Or", testOr, spec.Parallel())
	suite.Run(t)
}

func ContainerLogs(id string) func() string {
	docker := occam.NewDocker()

	return func() string {
		logs, _ := docker.Container.Logs.Execute(id)
		return logs.String()
	}
}
