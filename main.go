// Package main _
package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	version, createdTime, localVersion := "", "", ""

	wg := sync.WaitGroup{}
	wg.Add(3)

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

	go func() {
		defer wg.Done()
		var err error
		localVersion, err = getLocalVersion()
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()

	fmt.Printf("local version:  %s\n", localVersion)
	fmt.Printf("github version: %s (%s)\n", version, createdTime)

	if version != localVersion {
		answer := false
		prompt := &survey.Confirm{
			Message: "Do you want to update nvim?",
		}
		survey.AskOne(prompt, &answer)

		if answer {
			if err := update(version); err != nil {
				log.Fatal(err)
			}
			localVersion, err := getLocalVersion()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("new local version:  %s\n", localVersion)
		}
	}
}
