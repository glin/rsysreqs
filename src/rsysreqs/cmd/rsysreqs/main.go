package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"rsysreqs/rule"
)

func main() {
	const defaultSysreqs = "libXML2, curl; pkgA"
	const defaultRulesDir = "../rsysreqs-db/sysreqs/"

	sysreqs := flag.String("s", defaultSysreqs, "system requirements")
	rulesDir := flag.String("d", defaultRulesDir, "use rules from this directory")

	flag.Parse()

	rules, err := readRules(*rulesDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	matched, err := rule.MatchRules(*sysreqs, rules)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("found %d rules\n", len(matched))
	for _, r := range matched {
		fmt.Println(r)
	}
}

func readRules(path string) (rules []rule.Rule, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return rules, err
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(path + file.Name())
		if err != nil {
			return rules, err
		}

		r := rule.Rule{}
		err = json.Unmarshal(b, &r)
		if err != nil {
			return rules, err
		}
		rules = append(rules, r)
	}

	return rules, err
}
