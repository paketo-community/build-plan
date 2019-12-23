package noapp_test

import (
	"testing"

	"github.com/cloudfoundry/no-app-cnb/noapp"
	"github.com/cloudfoundry/packit"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		build packit.BuildFunc
	)

	it.Before(func() {
		build = noapp.Build()
	})

	it("sets the start command when only the runtime is used", func() {
		result, err := build(packit.BuildContext{})
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(packit.BuildResult{
			Processes: []packit.Process{
				{
					Type:    "web",
					Command: "/bin/bash",
					Direct:  false,
				},
			},
		}))
	})
}
