package main

import (
	"log"
	"path"

	"github.com/karrick/golf"
)

func main() {
	installDir = golf.StringP('d', "dir", "", "Factorio server install directory. (Default: ./factorio)")

	golf.Parse()

	createInstallDir()

	log.Printf("Created install directory: %s", *installDir)
	log.Print("Cloning factorio-init...")
	cloneRepo("factorio-init", path.Join(*installDir, "factorio-init"))
}
