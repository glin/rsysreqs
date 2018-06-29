package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/glin/rsysreqs"
)

func main() {
	const defaultSysreqs = "libXML2, curl; pkgA"
	const defaultRulesDir = "../rsysreqs-db/sysreqs/"

	sysreqs := flag.String("s", defaultSysreqs, "system requirements")
	rulesDir := flag.String("d", defaultRulesDir, "use rules from this directory")

	flag.Parse()

	rules, err := rsysreqs.ReadRules(*rulesDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	matched, err := rules.FindRules(*sysreqs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("found %d rules\n", len(matched))
	for _, r := range matched {
		fmt.Println(r)
	}
}
