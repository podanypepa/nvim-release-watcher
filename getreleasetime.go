package main

import (
	"github.com/gocolly/colly"
)

const (
	// AssetsPage is URL of web page with detail info about build files
	AssetsPage = "https://github.com/neovim/neovim/releases/expanded_assets/nightly"
)

func getReleaseFileTime() (string, error) {
	c := colly.NewCollector()
	createdTime := ""
	c.OnHTML("relative-time", func(e *colly.HTMLElement) {
		if createdTime == "" {
			createdTime = e.Text
		}
	})
	c.Visit(AssetsPage)
	return createdTime, nil
}
