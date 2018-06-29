package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/glin/rsysreqs"
)

func main() {
	sysreqs := flag.String("s", "", "system requirements")
	rulesDir := flag.String("d", "", "use rules from this directory")

	flag.Parse()

	if *sysreqs == "" || *rulesDir == "" {
		flag.Usage()
		os.Exit(2)
	}

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
