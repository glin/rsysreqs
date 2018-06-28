package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"rsysreqs/rule"

	"github.com/gin-gonic/gin"
)

func main() {
	const defaultRulesDir = "../rsysreqs-db/sysreqs/"
	rulesDir := flag.String("d", defaultRulesDir, "use rules from this directory")

	flag.Parse()

	rules, err := readRules(*rulesDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := gin.Default()
	r.GET("/rules", func(c *gin.Context) {
		sysreqs := c.DefaultQuery("sysreqs", "")

		matched, err := rule.MatchRules(sysreqs, rules)
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

func readRules(path string) (rules []rule.Rule, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return rules, err
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(path + file.Name())
		if err != nil {
			return rules, err
		}

		r := rule.Rule{}
		err = json.Unmarshal(b, &r)
		if err != nil {
			return rules, err
		}
		rules = append(rules, r)
	}

	return rules, err
}
