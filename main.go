// Package main _
package main

import (
	"fmt"
	"log"
	"sync"
)

const (
	// DevelopmentPage is URL of web page with basic prerelease build
	DevelopmentPage = "https://github.com/neovim/neovim/releases"
	// AssetsPage is URL of web page with detail info about build files
	AssetsPage = "https://github.com/neovim/neovim/releases/expanded_assets/nightly"
	// NvimFile is required release file
	NvimFile = "nvim-macos.tar.gz"
)

func main() {
	version, createdTime := "", ""

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		var err error
		version, err = getVersion()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		createdTime, err = getReleaseFileTime()
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()

	fmt.Println("version:    ", version)
	fmt.Println("createdTime:", createdTime)
}

func getVersion() (string, error) {
	return "", nil
}

func getReleaseFileTime() (string, error) {
	return "", nil
}
