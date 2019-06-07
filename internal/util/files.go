package util

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CopyFile(src, dst string) (err error) {
	srcContent, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0777)
	defer f.Close()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, srcContent, 0666)
	if err != nil {
		return err
	}
	return nil
}

func GetContentFile(path string) ([]string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	fmt.Println("[*] Reading NSRL DB")
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	content := lines
	fmt.Println("[*] Finished reading file: " + path)
	return content, nil
}

func GetHashFile(algo string, src io.Reader) (string, error) {
	var h hash.Hash
	switch strings.ToLower(algo) {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "crc32":
		h = crc32.New(crc32.IEEETable)
	}
	_, err := io.Copy(h, src)
	if err != nil {
		return "", err
	}

	hashFile := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
	return hashFile, nil
}
