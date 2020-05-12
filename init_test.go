package main_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitBuildPlan(t *testing.T) {
	suite := spec.New("build-plan", spec.Report(report.Terminal{}))
	suite("Build", testBuild)
	suite("Detect", testDetect)
	suite("EnvironmentParser", testEnvironmentParser)
	suite.Run(t)
}
