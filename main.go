// Package main _
package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	version, createdTime, localVersion := "", "", ""

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		var err error
		version, err = getVersion()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get version")
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		createdTime, err = getReleaseFileTime()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get release file time")
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		localVersion, err = getLocalVersion()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get local version")
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
				log.Fatal().Err(err).Msg("failed to update")
			}
			localVersion, err := getLocalVersion()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to get local version")
			}
			fmt.Printf("new local version:  %s\n", localVersion)
		}
	}
}
