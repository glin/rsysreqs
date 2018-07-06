package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"rsysreqs/rules"
	"rsysreqs/scripts"
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

	readRules, err := rules.ReadRules(*rulesDir)
	if err != nil {
		fmt.Println("error reading rules:", err)
		os.Exit(1)
	}

	matched, err := readRules.FindRules(*sysreqs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	system := rules.System{
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

	generator, err := scripts.NewScriptGenerator(system)
	if err != nil {
		fmt.Println("error generating install scripts:", err)
		os.Exit(1)
	}

	installScripts := generator.InstallScripts(packages)

	pkgActions := struct {
		Packages       []string `json:"packages"`
		InstallScripts []string `json:"installScripts"`
	}{
		packages,
		installScripts,
	}

	pkgActionsJson, err := json.MarshalIndent(pkgActions, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", pkgActionsJson)
}
