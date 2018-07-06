package scripts

import (
	"errors"

	"rsysreqs/rules"
)

var ErrUnsupportedSystem = errors.New("unsupported system")

type ScriptGenerator interface {
	InstallScripts(packages []string) []string
}

func NewScriptGenerator(sys rules.System) (ScriptGenerator, error) {
	switch {
	case sys.Os == "linux" && sys.Distribution == "ubuntu":
		return ubuntuScriptGenerator{}, nil
	}
	return nil, ErrUnsupportedSystem
}
