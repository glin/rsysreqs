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
	Architecture string `json:"architecture,omitempty"`
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
