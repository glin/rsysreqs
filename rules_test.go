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

func (s *Suite) TestFindPackages(c *check.C) {
	rulePkgA := Rule{
		Sysreqs: []string{"\\bpkgA\\b"},
		Dependencies: []Dependency{
			{
				Packages:    []string{"pkgA-1", "pkgA-2", "pkgC"},
				Constraints: []Constraint{{Distribution: "ubuntu"}},
			},
		},
	}

	rulePkgB := Rule{
		Sysreqs: []string{"\\bpkgB\\b"},
		Dependencies: []Dependency{
			{
				Packages:    []string{"pkgB-1", "pkgB-2", "pkgC"},
				Constraints: []Constraint{{Distribution: "ubuntu", Architecture: "i386"}},
			},
		},
	}

	rules := Rules{rulePkgA, rulePkgB}

	found, err := rules.FindRules("pkga, pkgb")
	c.Assert(err, check.IsNil)

	packages, err := found.FindPackages(System{Distribution: "ubuntu"})
	c.Assert(err, check.IsNil)
	c.Assert(packages, check.DeepEquals, []string{"pkgA-1", "pkgA-2", "pkgC", "pkgB-1", "pkgB-2"})
}
