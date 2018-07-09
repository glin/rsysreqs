package scripts

import "fmt"

type ubuntuScriptGenerator struct{}

func (g ubuntuScriptGenerator) InstallScripts(packages []string) []string {
	var scripts []string
	for _, pkg := range packages {
		script := fmt.Sprintf("apt-get install -y %s", pkg)
		scripts = append(scripts, script)
	}
	return scripts
}
