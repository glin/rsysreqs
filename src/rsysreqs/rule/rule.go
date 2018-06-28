package rule

import (
	"errors"
	"regexp"
)

var ErrNoMatchingRules = errors.New("no matching rules found")

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

func MatchRules(sysreqs string, rules []Rule) (matched []Rule, err error) {
	for _, rule := range rules {
		for _, pattern := range rule.Sysreqs {
			// TODO check for existing flags
			match, _ := regexp.MatchString("(?i)"+pattern, sysreqs)
			if match {
				matched = append(matched, rule)
			}
		}
	}

	if len(matched) == 0 {
		err = ErrNoMatchingRules
	}

	return matched, err
}
