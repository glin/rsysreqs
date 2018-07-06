package scripts

import (
	"errors"
	"fmt"

	"rsysreqs/rules"
)

var ErrUnsupportedSystem = errors.New("unsupported system")

func GenerateInstallScripts(sys rules.System, packages []string) ([]string, error) {
	var cmds []string
	for _, pkg := range packages {
		cmd, err := generateInstallScript(sys, pkg)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	}

	return cmds, nil
}

func generateInstallScript(sys rules.System, pkg string) (string, error) {
	var script string
	switch {
	case sys.Os == "linux" && sys.Distribution == "ubuntu":
		script = ubuntuInstallScript(pkg)
	}

	if script == "" {
		return "", ErrUnsupportedSystem
	}

	return script, nil
}

func ubuntuInstallScript(pkg string) string {
	return fmt.Sprintf("apt-get install -y %s", pkg)
}
