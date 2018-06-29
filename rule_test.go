package rsysreqs

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type Suite struct{}

var _ = check.Suite(&Suite{})

func (s *Suite) TestRuleMatch(c *check.C) {
	sysreqs := "pkga, pkgb"

	rulePkgA := Rule{
		Sysreqs: []string{"\\bpkgA\\b"},
		Dependencies: []Dependency{
			{
				Packages:    []string{"pkgA"},
				Constraints: []Constraint{{Distribution: "ubuntu"}},
			},
		},
	}

	matched, err := rulePkgA.Match(sysreqs)

	c.Assert(err, check.IsNil)
	c.Assert(matched, check.Equals, true)
}
