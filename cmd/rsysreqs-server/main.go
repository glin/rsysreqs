package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/glin/rsysreqs"
)

func main() {
	rulesDir := flag.String("d", "", "use rules from this directory")

	flag.Parse()

	if *rulesDir == "" {
		flag.Usage()
		os.Exit(2)
	}

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
