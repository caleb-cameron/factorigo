package main

import (
	"log"

	"github.com/karrick/golf"
)

func main() {
	installDir = golf.StringP('d', "dir", "", "Factorio server install directory. (Default: ./factorio)")

	golf.Parse()

	createInstallDir()

	log.Printf("Created install directory: %s", *installDir)
}
