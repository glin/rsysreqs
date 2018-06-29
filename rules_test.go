package rsysreqs

import (
	"gopkg.in/check.v1"
)

func (s *Suite) TestFindRules(c *check.C) {
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

	rules := Rules{rulePkgA}

	found, err := rules.FindRules(sysreqs)

	c.Assert(err, check.IsNil)
	c.Assert(found, check.HasLen, 1)
	c.Assert(found[0], check.DeepEquals, rulePkgA)
}
