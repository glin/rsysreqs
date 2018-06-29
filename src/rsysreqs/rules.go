package rsysreqs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var ErrNoMatchingRules = errors.New("no matching rules found")

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

func ReadRules(path string) (rules Rules, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return rules, err
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(path + file.Name())
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
