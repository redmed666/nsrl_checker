package main

import (
	"io/ioutil"
	"os"

	nsrl "github.com/redmed666/nsrl_checker/internal/nsrl"
	util "github.com/redmed666/nsrl_checker/internal/util"

	"github.com/akamensky/argparse"
)

var (
	nsrlFile      *string
	outDir        *string
	exeDir        *string
	hashAlgo      *string
	contentNsrl   map[string]int
	exeDirInfo    []os.FileInfo
	dirs          util.Dir
	nmbThreadsArg *int
)

func ParseArgs() {
	argp := argparse.NewParser("nsrl_checker", "Checks folder of exes against NSRL hashes")
	nsrlFile = argp.String("n", "nsrl", &argparse.Options{Required: true, Help: "Path to the NSRL DB file"})
	outDir = argp.String("o", "outdir", &argparse.Options{Required: false, Help: "Output directory", Default: "."})
	exeDir = argp.String("e", "exedir", &argparse.Options{Required: true, Help: "Folder containing the exes"})
	hashAlgo = argp.Selector("H", "hash", []string{"md5", "sha1", "crc32"}, &argparse.Options{Required: false, Help: "Hash algorithm used for the comparison", Default: "sha1"})
	nmbThreadsArg = argp.Int("t", "threads", &argparse.Options{Required: false, Help: "Max number of threads", Default: 10})

	err := argp.Parse(os.Args)
	util.Check(err)
}

func init() {
	ParseArgs()
	var err error
	f, err := os.Open(*nsrlFile)
	util.Check(err)

	contentNsrl = nsrl.ParseNSRLHashFile(f, *hashAlgo)

	exeDirInfo, err = ioutil.ReadDir(*exeDir)
	util.Check(err)

	dirs, err = util.GetDirs(*outDir)
	util.Check(err)

	err = util.CreateDirs(dirs)
	util.Check(err)
}

func main() {
	var nmbThreads int
	if *nmbThreadsArg > len(exeDirInfo) {
		nmbThreads = len(exeDirInfo)
	} else {
		nmbThreads = *nmbThreadsArg
	}

	guard := make(chan struct{}, nmbThreads)
	for _, file := range exeDirInfo {
		exePath := *exeDir + "/" + file.Name()
		guard <- struct{}{}
		go util.SortExe(exePath, &contentNsrl, *hashAlgo, dirs, guard)
	}
}
