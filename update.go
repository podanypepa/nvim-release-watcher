package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	nvimTarUrl  = "https://github.com/neovim/neovim/releases/download/nightly/nvim-macos-arm64.tar.gz"
	dstDir      = "/Users/pepa/bin/neovim/"
	symlinkName = "/Users/pepa/bin/nvim"
)

func update(version string) error {
	fmt.Printf("downloading %s\n", nvimTarUrl)

	version = strings.Split(version, " ")[1]
	newDir := path.Join(dstDir, version)

	if _, err := os.Stat(newDir); os.IsExist(err) {
		log.Error().Caller().Msgf("dir %s is already exists", newDir)
		return fmt.Errorf("dir %s is already exists", newDir)
	}

	resp, err := http.Get(nvimTarUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := extractTarGz(resp.Body); err != nil {
		return err
	}

	if err := os.Chmod("./nvim-macos-arm64/bin/nvim", 0755); err != nil {
		log.Fatal().Caller().Err(err).Msg("failed to chmod")
	}

	if err := os.Rename("./nvim-macos-arm64", newDir); err != nil {
		return err
	}

	if err := os.Remove(symlinkName); err != nil {
		return err
	}

	srcBin := path.Join(newDir, "bin/nvim")
	_ = os.Symlink(srcBin, symlinkName)

	return os.RemoveAll("./nvim-macos-arm64")
}

func extractTarGz(gzipStream io.Reader) error {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal().Caller().Err(err).Msg("failed to read tar")
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(header.Name, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
			outFile.Close()

		default:
			return fmt.Errorf("error: %b, %s", header.Typeflag, header.Name)
		}

	}
	return nil
}
