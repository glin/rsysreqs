package rsysreqs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
)

var (
	ErrNoMatchingRules = errors.New("no matching rules found")
	ErrNoPackages      = errors.New("no packages found")
)

type Rules []Rule

func (rules Rules) FindRules(sysreqs string) (found Rules, err error) {
	for _, rule := range rules {
		matched, err := rule.Match(sysreqs)
		if err != nil {
			return rules, err
		}
		if matched {
			found = append(found, rule)
		}
	}

	if len(found) == 0 {
		err = ErrNoMatchingRules
	}

	return found, err
}

func (rules Rules) FindPackages(system System) (packages []string, err error) {
	seen := map[string]bool{}

	for _, rule := range rules {
		found := rule.FindPackages(system)
		for _, pkg := range found {
			if !seen[pkg] {
				seen[pkg] = true
				packages = append(packages, pkg)
			}
		}
	}

	if len(packages) == 0 {
		err = ErrNoPackages
	}

	return packages, err
}

func ReadRules(dirname string) (rules Rules, err error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return rules, err
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(filepath.Join(dirname, file.Name()))
		if err != nil {
			return rules, err
		}

		r := Rule{}
		err = json.Unmarshal(b, &r)
		if err != nil {
			return rules, err
		}
		rules = append(rules, r)
	}

	return rules, err
}
