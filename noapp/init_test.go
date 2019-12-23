package noapp_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitNode(t *testing.T) {
	suite := spec.New("noapp", spec.Report(report.Terminal{}))
	suite("Build", testBuild)
	suite("Detect", testDetect)
	suite("EnvironmentParser", testEnvironmentParser)
	suite.Run(t)
}
