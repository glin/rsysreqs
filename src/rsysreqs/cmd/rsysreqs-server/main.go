package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"rsysreqs/rules"
	"rsysreqs/scripts"
)

var ErrMissingParams = errors.New("missing required parameters")

func main() {
	rulesDir := flag.String("d", "", "use rules from this directory")

	flag.Parse()

	if *rulesDir == "" {
		flag.Usage()
		os.Exit(2)
	}

	r := gin.Default()

	r.Use(contextRules(*rulesDir))

	r.GET("/packages", getPackages)

	r.Run()
}

func contextRules(rulesDir string) gin.HandlerFunc {
	readRules, err := rules.ReadRules(rulesDir)

	if err != nil {
		log.Fatalf("error reading rules: %s", err)
	}

	return func(c *gin.Context) {
		c.Set("rules", readRules)
		c.Next()
	}
}

func getPackages(c *gin.Context) {
	r, exists := c.Get("rules")
	ctxRules, ok := r.(rules.Rules)
	if !exists || !ok {
		log.Fatalf("missing or invalid rules")
	}

	sysreqs := c.Query("sysreqs")
	sysOs := c.Query("os")
	sysDistribution := c.Query("dist")
	sysRelease := c.Query("release")
	sysArch := c.Query("arch")

	if sysreqs == "" || sysOs == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrMissingParams.Error(),
		})
		return
	}

	matched, err := ctxRules.FindRules(sysreqs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	system := rules.System{
		Os:           sysOs,
		Distribution: sysDistribution,
		Release:      sysRelease,
		Architecture: sysArch,
	}

	packages, err := matched.FindPackages(system)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	generator, err := scripts.NewScriptGenerator(system)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	installScripts := generator.InstallScripts(packages)

	c.JSON(http.StatusOK, gin.H{
		"packages":        packages,
		"install_scripts": installScripts,
	})
}
