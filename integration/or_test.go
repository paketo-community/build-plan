package integration_test

import (
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

func testOr(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		Eventually = NewWithT(t).Eventually

		pack   occam.Pack
		docker occam.Docker

		image     occam.Image
		container occam.Container
		name      string
	)

	it.Before(func() {
		pack = occam.NewPack().WithVerbose()
		docker = occam.NewDocker()

		var err error
		name, err = occam.RandomName()
		Expect(err).NotTo(HaveOccurred())
	})

	it.After(func() {
		Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
		Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
		Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
	})

	it("should build a working OCI image for with the python runtime", func() {
		var err error
		image, _, err = pack.WithNoColor().Build.
			WithPullPolicy("never").
			WithBuildpacks(cpythonBuildpack, buildpack).
			Execute(name, filepath.Join("testdata", "or"))
		Expect(err).ToNot(HaveOccurred())

		container, err = docker.Container.Run.
			WithEnv(map[string]string{"PORT": "8080"}).
			WithPublish("8080").
			WithPublishAll().
			WithCommand("python --version && python -m http.server $PORT").
			Execute(image.ID)
		Expect(err).NotTo(HaveOccurred())

		Eventually(container).Should(BeAvailable(), ContainerLogs(container.ID))

		logs, err := docker.Container.Logs.Execute(container.ID)
		Expect(err).NotTo(HaveOccurred())

		Expect(logs).To(MatchRegexp(`Python \d+\.\d+\.\d+`))
	})
}
