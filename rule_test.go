package rsysreqs

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type Suite struct{}

var _ = check.Suite(&Suite{})

func (s *Suite) TestRuleMatch(c *check.C) {
	rulePkgA := Rule{
		Sysreqs: []string{"\\bpkgA\\b"},
		Dependencies: []Dependency{
			{
				Packages:    []string{"pkgA"},
				Constraints: []Constraint{{Distribution: "ubuntu"}},
			},
		},
	}

	tests := []struct {
		sysreqs         string
		expectedMatched bool
	}{
		{"pkgA", true},
		{"pkga, pkgb", true},
		{"pkgb;pkga", true},
		{"pkga pkgb", true},
		{"pkgb", false},
		{"pkgAB", false},
	}

	for _, test := range tests {
		matched, err := rulePkgA.Match(test.sysreqs)
		c.Assert(err, check.IsNil)
		c.Assert(matched, check.Equals, test.expectedMatched)
	}
}
