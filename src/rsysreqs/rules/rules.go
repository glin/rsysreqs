package rules

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

func (rules Rules) FindRules(sysreqs string) (Rules, error) {
	var found Rules
	for _, rule := range rules {
		matched, err := rule.Match(sysreqs)
		if err != nil {
			return nil, err
		}
		if matched {
			found = append(found, rule)
		}
	}

	if len(found) == 0 {
		return nil, ErrNoMatchingRules
	}

	return found, nil
}

func (rules Rules) FindPackages(system System) ([]string, error) {
	var packages []string
	seen := make(map[string]bool, 0)
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
		return nil, ErrNoPackages
	}

	return packages, nil
}

func ReadRules(dirname string) (Rules, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	var rules Rules
	for _, file := range files {
		b, err := ioutil.ReadFile(filepath.Join(dirname, file.Name()))
		if err != nil {
			return nil, err
		}

		r := Rule{}
		err = json.Unmarshal(b, &r)
		if err != nil {
			return nil, err
		}
		rules = append(rules, r)
	}

	return rules, err
}
