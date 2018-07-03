package rsysreqs

import (
	"regexp"
)

type Rule struct {
	Sysreqs      []string     `json:"sysreqs"`
	Dependencies []Dependency `json:"dependencies"`
}

type Dependency struct {
	Packages    []string     `json:"packages"`
	Constraints []Constraint `json:"constraints"`
}

type Constraint struct {
	Os           string `json:"os,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	ReleaseMin   string `json:"releaseMin,omitempty"`
	ReleaseMax   string `json:"releaseMax,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

type System struct {
	Os           string
	Distribution string
	Release      string
	Architecture string
}

func (rule Rule) Match(sysreqs string) (matched bool, err error) {
	for _, pattern := range rule.Sysreqs {
		// TODO check for existing flags
		matched, err = regexp.MatchString("(?i)"+pattern, sysreqs)
		if matched || err != nil {
			break
		}
	}
	return matched, err
}

func (rule Rule) FindPackages(sys System) (packages []string) {
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
