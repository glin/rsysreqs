package rule

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type Suite struct{}

var _ = check.Suite(&Suite{})

func (s *Suite) TestMatchRules(c *check.C) {
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

	rules := []Rule{rulePkgA}

	matched, err := MatchRules(sysreqs, rules)

	c.Assert(err, check.IsNil)
	c.Assert(matched, check.HasLen, 1)
	c.Assert(matched[0], check.DeepEquals, rulePkgA)
}
