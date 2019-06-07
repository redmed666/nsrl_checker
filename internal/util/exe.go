package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/radare/r2pipe-go"
)

func IsSigned(exePath string, f *os.File) (bool, error) {
	r2, err := r2pipe.NewPipe(exePath)
	result := false
	defer r2.Close()
	if err != nil {
		return result, err
	} else {
		infosRaw, err := r2.Cmdj("ij")
		if err != nil {
			return result, err
		} else {
			infos := infosRaw.(map[string]interface{})
			if infos["bin"] != nil {
				infosBin := infos["bin"].(map[string]interface{})
				if infosBin["signed"] != nil {
					infosSigned := infosBin["signed"].(bool)
					if infosSigned == true {
						result = true
					} else {
						result = false
					}
				} else {
					result = false
				}
				if err != nil {
					return result, err
				}
			}
		}
	}
	return result, err
}

func SortExe(exePath string, contentNsrl *map[string]int, hashAlgo string, dirs Dir, guard chan struct{}) {
	f, err := os.Open(exePath)
	defer f.Close()
	Check(err)

	hashFile, err := GetHashFile(hashAlgo, f)
	Check(err)

	if (*contentNsrl)[hashFile] == 1 {
		out := dirs.OutSafe + "/" + filepath.Base(f.Name())
		fmt.Printf("[*] Copying %s to %s\n", exePath, out)
		err = CopyFile(exePath, out)
	} else {
		signed, err := IsSigned(exePath, f)
		Check(err)

		if signed == true {
			out := dirs.OutUnsafeSigned + "/" + filepath.Base(f.Name())
			fmt.Printf("[*] Copying %s to %s\n", exePath, out)
			err = CopyFile(exePath, out)
		} else {
			out := dirs.OutUnsafeUnsigned + "/" + filepath.Base(f.Name())
			fmt.Printf("[*] Copying %s to %s\n", exePath, out)
			err = CopyFile(exePath, out)
		}
	}
	Check(err)
	<-guard
}
