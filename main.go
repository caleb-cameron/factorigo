package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/karrick/golf"
)

func main() {
	var err error

	installDir := golf.StringP('d', "dir", "", "Factorio server install directory. (Default: ./factorio)")

	golf.Parse()

	if *installDir == "" {
		workingDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Could not get working dir: %s", err)
		}
		*installDir = path.Join(workingDir, "factorio")
	}

	err = setupInstallDir(*installDir)

	if err != nil {
		log.Fatalf(fmt.Sprintf("Cannot setup install directory %s: %s", *installDir, err))
	}

	log.Printf("Created install directory: %s", *installDir)
}
