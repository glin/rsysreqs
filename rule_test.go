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
		sysreqs  string
		expected bool
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
		c.Assert(matched, check.Equals, test.expected)
	}
}

func (s *Suite) TestRuleFindPackages(c *check.C) {
	rule := Rule{
		Sysreqs: []string{"\\bpkgA\\b"},
		Dependencies: []Dependency{
			{
				Packages:    []string{"pkgA"},
				Constraints: []Constraint{{Distribution: "ubuntu"}},
			},
			{
				Packages:    []string{"pkgA-1", "pkgA-2"},
				Constraints: []Constraint{{Os: "linux"}},
			},
			{
				Packages:    []string{"pkgA-centos"},
				Constraints: []Constraint{{Distribution: "centos"}},
			},
		},
	}

	system := System{Distribution: "ubuntu", Os: "linux"}
	packages := rule.FindPackages(system)
	c.Assert(packages, check.DeepEquals, []string{"pkgA", "pkgA-1", "pkgA-2"})
}

func (s *Suite) TestConstraintSatisfiedBy(c *check.C) {
	tests := []struct {
		system     System
		constraint Constraint
		expected   bool
	}{
		{System{Distribution: "ubuntu"}, Constraint{Distribution: "ubuntu"}, true},
		{System{Distribution: "ubuntu"}, Constraint{Distribution: "centos"}, false},
		{System{Os: "windows"}, Constraint{Os: "windows"}, true},
		{System{Os: "windows"}, Constraint{Os: "linux"}, false},
		{System{Architecture: "i386"}, Constraint{Architecture: "i386"}, true},
		{System{Architecture: "i386"}, Constraint{Architecture: "amd64"}, false},
		{System{Release: "14.04"}, Constraint{ReleaseMin: "14.04", ReleaseMax: "16.04"}, true},
		{System{Release: "14.04"}, Constraint{ReleaseMin: "16.04"}, false},
		{System{Release: "14.04"}, Constraint{ReleaseMin: "14.04", ReleaseMax: "14.04"}, true},
		{System{Release: "14.04"}, Constraint{ReleaseMax: "12.04"}, false},
		{System{Os: "linux", Distribution: "ubuntu"}, Constraint{Distribution: "ubuntu"}, true},
		{System{Os: "linux", Distribution: "ubuntu"}, Constraint{Os: "windows", Distribution: "ubuntu"}, false},
	}

	for _, test := range tests {
		system, constraint := test.system, test.constraint
		satisfied := constraint.satisfiedBy(system)
		c.Assert(satisfied, check.Equals, test.expected,
			check.Commentf("system: %s, constraint: %s, expected: %v", system, constraint, test.expected))
	}
}
