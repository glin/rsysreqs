package rule

import (
	"errors"
	"regexp"
)

var ErrNoMatchingRules = errors.New("no matching rules found")

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
	Architecture string
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
