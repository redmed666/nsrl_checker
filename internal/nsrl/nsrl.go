package nsrl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ParseNSRLHashFile Parse column of NSRLFile.txt and output an unordered hashmap
func ParseNSRLHashFile(f io.Reader, algo string) map[string]int {
	/*
		First line is useless
		Header:
		SHA-1,MD5,CRC32,FileName,FileSize,ProductCode,OpSystemCode,SpecialCode
	*/
	fmt.Println("[*] Parsing NSRL file")
	results := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Scan() // skipping the first line == header
	for scanner.Scan() {
		line := scanner.Text()
		lineSplitted := strings.Split(line, ",")
		var rawResult string

		switch algo {
		case "sha1":
			rawResult = lineSplitted[0]
		case "md5":
			rawResult = lineSplitted[1]
		case "crc32":
			rawResult = lineSplitted[2]
		}

		result := strings.ToLower(strings.Replace(rawResult, "\"", "", -1))
		results[result] = 1
	}

	return results
}
