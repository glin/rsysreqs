package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/glin/rsysreqs"
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

	r.Use(rules(*rulesDir))

	r.GET("/packages", getPackages)

	r.Run()
}

func rules(rulesDir string) gin.HandlerFunc {
	rules, err := rsysreqs.ReadRules(rulesDir)

	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		c.Set("rules", rules)
		c.Next()
	}
}

func getPackages(c *gin.Context) {
	r, exists := c.Get("rules")
	rules, ok := r.(rsysreqs.Rules)
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

	matched, err := rules.FindRules(sysreqs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	system := rsysreqs.System{
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

	c.JSON(http.StatusOK, gin.H{
		"packages": packages,
	})
}
