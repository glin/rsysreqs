package rules

import (
	"regexp"
)

type Rule struct {
	Sysreqs      []string
	Dependencies []Dependency
}

type Dependency struct {
	Packages    []string
	Constraints []Constraint
}

type Constraint struct {
	Os           string
	Distribution string
	ReleaseMin   string
	ReleaseMax   string
	Architecture string
}

type System struct {
	Os           string
	Distribution string
	Release      string
	Architecture string
}

func (rule Rule) Match(sysreqs string) (bool, error) {
	for _, pattern := range rule.Sysreqs {
		// TODO check for existing flags
		matched, err := regexp.MatchString("(?i)"+pattern, sysreqs)
		if matched || err != nil {
			return matched, err
		}
	}
	return false, nil
}

func (rule Rule) FindPackages(sys System) []string {
	var packages []string
	for _, dep := range rule.Dependencies {
		for _, constraint := range dep.Constraints {
			if constraint.satisfiedBy(sys) {
				packages = append(packages, dep.Packages...)
			}
		}
	}
	return packages
}

func (c Constraint) satisfiedBy(sys System) bool {
	if c.Os != "" && sys.Os != "" && sys.Os != c.Os {
		return false
	}
	if c.Distribution != "" && sys.Distribution != "" && sys.Distribution != c.Distribution {
		return false
	}
	if c.ReleaseMin != "" && sys.Release != "" && sys.Release < c.ReleaseMin {
		return false
	}
	if c.ReleaseMax != "" && sys.Release != "" && sys.Release > c.ReleaseMax {
		return false
	}
	if c.Architecture != "" && sys.Architecture != "" && sys.Architecture != c.Architecture {
		return false
	}
	return true
}
