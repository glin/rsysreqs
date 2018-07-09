package scripts

import (
	"testing"

	"gopkg.in/check.v1"

	"rsysreqs/rules"
)

func Test(t *testing.T) { check.TestingT(t) }

type Suite struct{}

var _ = check.Suite(&Suite{})

func (s *Suite) TestNewScriptGenerator(c *check.C) {
	sys := rules.System{Os: "linux", Distribution: "ubuntu"}
	generator, err := NewScriptGenerator(sys)
	c.Assert(err, check.IsNil)
	c.Assert(generator, check.FitsTypeOf, ubuntuScriptGenerator{})
}

func (s *Suite) TestNewScriptGeneratorUnsupported(c *check.C) {
	sys := rules.System{Os: "unsupported_os"}
	generator, err := NewScriptGenerator(sys)
	c.Assert(err, check.Equals, ErrUnsupportedSystem)
	c.Assert(generator, check.IsNil)
}
