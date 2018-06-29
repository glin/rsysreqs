package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"rsysreqs"
)

func main() {
	const defaultRulesDir = "../rsysreqs-db/sysreqs/"
	rulesDir := flag.String("d", defaultRulesDir, "use rules from this directory")

	flag.Parse()

	rules, err := rsysreqs.ReadRules(*rulesDir)

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/rules", func(c *gin.Context) {
		sysreqs := c.DefaultQuery("sysreqs", "")

		matched, err := rules.FindRules(sysreqs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"rules": matched,
		})
	})

	r.Run()
}
