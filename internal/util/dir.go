package util

import (
	"errors"
	"os"
)

type Dir struct {
	Out               string
	OutSafe           string
	OutUnsafe         string
	OutUnsafeSigned   string
	OutUnsafeUnsigned string
}

func GetDirs(outDir string) (Dir, error) {
	var dirs Dir
	if outDir == "" {
		err := errors.New("outDir is not defined")
		return dirs, err
	}

	dirs.Out = outDir
	dirs.OutSafe = outDir + "/safe"
	dirs.OutUnsafe = outDir + "/unsafe"
	dirs.OutUnsafeSigned = outDir + "/unsafe/signed"
	dirs.OutUnsafeUnsigned = outDir + "/unsafe/unsigned"
	return dirs, nil
}

func CreateDirs(dirs Dir) error {

	if _, err := os.Stat(dirs.Out); os.IsNotExist(err) {
		err = os.Mkdir(dirs.Out, 0777)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(dirs.OutSafe); os.IsNotExist(err) {
		err = os.Mkdir(dirs.OutSafe, 0777)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(dirs.OutUnsafe); os.IsNotExist(err) {
		err = os.Mkdir(dirs.OutUnsafe, 0777)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(dirs.OutUnsafeSigned); os.IsNotExist(err) {
		err = os.Mkdir(dirs.OutUnsafeSigned, 0777)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(dirs.OutUnsafeUnsigned); os.IsNotExist(err) {
		err = os.Mkdir(dirs.OutUnsafeUnsigned, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
