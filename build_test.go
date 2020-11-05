package buildplan_test

import (
	"testing"

	"github.com/paketo-buildpacks/packit"
	buildplan "github.com/paketo-community/build-plan"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		build packit.BuildFunc
	)

	it.Before(func() {
		build = buildplan.Build()
	})

	it("sets the start command when only the runtime is used", func() {
		result, err := build(packit.BuildContext{})
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(packit.BuildResult{}))
	})
}
