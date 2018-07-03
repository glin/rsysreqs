package main

import (
	"flag"
	"fmt"
	"os"

	"encoding/json"

	"github.com/glin/rsysreqs"
)

func main() {
	sysreqs := flag.String("s", "", "system requirements")
	rulesDir := flag.String("d", "", "use rules from this directory")
	sysOs := flag.String("os", "", "operating system")
	sysDistribution := flag.String("dist", "", "distribution")
	sysRelease := flag.String("release", "", "release")
	sysArch := flag.String("arch", "", "architecture")

	flag.Parse()

	if *sysreqs == "" || *rulesDir == "" || *sysOs == "" {
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

	system := rsysreqs.System{
		Os:           *sysOs,
		Distribution: *sysDistribution,
		Release:      *sysRelease,
		Architecture: *sysArch,
	}

	packages, err := matched.FindPackages(system)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	packagesJson, err := json.MarshalIndent(packages, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", packagesJson)
}
