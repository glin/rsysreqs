package scripts

import (
	"testing"

	"gopkg.in/check.v1"

	"rsysreqs/rules"
)

func Test(t *testing.T) { check.TestingT(t) }

type Suite struct{}

var _ = check.Suite(&Suite{})

func (s *Suite) TestGenerateInstallScripts(c *check.C) {
	sys := rules.System{Os: "linux", Distribution: "ubuntu"}
	packages := []string{"pkgA", "pkgB"}
	scripts, err := GenerateInstallScripts(sys, packages)

	c.Assert(err, check.IsNil)
	c.Assert(scripts, check.HasLen, 2)
	c.Assert(scripts[0], check.DeepEquals, "apt-get install -y pkgA")
	c.Assert(scripts[1], check.DeepEquals, "apt-get install -y pkgB")
}

func (s *Suite) TestGenerateInstallScript(c *check.C) {
	sys := rules.System{Os: "linux", Distribution: "ubuntu"}
	script, err := generateInstallScript(sys, "pkgA")

	c.Assert(err, check.IsNil)
	c.Assert(script, check.Equals, "apt-get install -y pkgA")
}
