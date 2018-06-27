package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	const defaultSysreqs = "libXML2, curl; pkgA"
	const defaultRulesDir = "../rsysreqs-db/sysreqs/"

	sysreqs := flag.String("s", defaultSysreqs, "system requirements")
	rulesDir := flag.String("d", defaultRulesDir, "rules directory")

	flag.Parse()

	rules, err := ReadRules(*rulesDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	matched, err := MatchRules(*sysreqs, rules)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("found %d rules\n", len(matched))
	for _, rule := range matched {
		fmt.Println(rule)
	}
}

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

func ReadRules(path string) (rules []Rule, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return rules, err
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(path + file.Name())
		if err != nil {
			return rules, err
		}

		rule := Rule{}
		err = json.Unmarshal(b, &rule)
		if err != nil {
			return rules, err
		}
		rules = append(rules, rule)
	}

	return rules, err
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
