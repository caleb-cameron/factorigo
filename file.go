package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

const (
	// rwxrwxr--
	INSTALL_DIR_PERMS = 0774
)

/*
	First check if the dir already exists. If not, create it.
	Then check that it's both writable and empty.
*/
func setupInstallDir(installDir string) error {
	exists, err := checkDirExists(installDir)

	if err != nil {
		return errors.New(fmt.Sprintf("Error checking if install directory %s exists: %s", installDir, err))
	}

	if !exists {
		err = os.MkdirAll(installDir, INSTALL_DIR_PERMS)
		if err != nil {
			return errors.New(fmt.Sprintf("Error creating install directory %s: %s", installDir, err))
		}
	}

	err = checkDirForWriting(installDir)
	if err != nil {
		return errors.New(fmt.Sprintf("Error checking if install dir %s is writable: %s", installDir, err))
	}

	isEmpty, err := checkDirIsEmpty(installDir)
	if err != nil {
		return errors.New(fmt.Sprintf("Error checking if install dir %s is empty: %s", installDir, err))
	}

	if !isEmpty {
		return errors.New(fmt.Sprintf("Install directory is not empty: %s", installDir))
	}

	return nil
}

func checkDirForWriting(dir string) error {
	return unix.Access(dir, unix.W_OK)
}

func checkDirIsEmpty(dir string) (bool, error) {
	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err // Either not empty or error, suits both cases
}

func checkDirExists(dir string) (bool, error) {
	f, err := os.Stat(dir)

	if err == nil && f.IsDir() {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func checkFileExists(dir string) (bool, error) {
	f, err := os.Stat(dir)

	if err == nil && !f.IsDir() {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
