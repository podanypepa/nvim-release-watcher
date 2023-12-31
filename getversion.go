package main

import (
	"strings"

	"github.com/gocolly/colly"
)

const (
	// DevelopmentPage is URL of web page with basic prerelease build
	DevelopmentPage = "https://github.com/neovim/neovim/releases"
)

func getVersion() (string, error) {
	c := colly.NewCollector()
	version := ""
	c.OnHTML("code", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "NVIM") && version == "" {
			version = e.Text
			info := strings.Split(e.Text, "\n")
			version = info[0]
		}
	})
	if err := c.Visit(DevelopmentPage); err != nil {
		return "", err
	}
	return version, nil
}
