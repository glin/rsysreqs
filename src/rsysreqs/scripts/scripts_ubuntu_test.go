package scripts

import (
	"gopkg.in/check.v1"
)

func (s *Suite) TestUbuntuInstallScripts(c *check.C) {
	generator := ubuntuScriptGenerator{}
	scripts := generator.InstallScripts([]string{"pkgA", "pkgB"})
	c.Assert(scripts, check.DeepEquals, []string{
		"apt-get install -y pkgA",
		"apt-get install -y pkgB",
	})
}
