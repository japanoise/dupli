package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var (
	scanner *bufio.Scanner
)

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	var files map[string][]string
	if len(os.Args) == 1 {
		path := "."
		files = iterateAndHash(path)
	} else {
		files = iterateAndHash(os.Args[1])
		if len(os.Args) > 2 {
			for _, path := range os.Args[2:] {
				morefiles := iterateAndHash(path)
				for hash, file := range morefiles {
					if files[hash] == nil {
						files[hash] = file
					} else {
						files[hash] = append(files[hash], file...)
					}
				}
			}
		}
	}
	for hash, file := range files {
		if len(file) > 1 {
			fmt.Println("hash", hash)
			delQuery(file)
		}
	}
}

func iterateAndHash(searchpath string) map[string][]string {
	ret := map[string][]string{}
	err := filepath.Walk(searchpath, func(path string, f os.FileInfo, err error) error {
		if isDir(path) {
			return nil
		}
		hash, err := hashFile(path)
		if err != nil {
			return err
		}
		if ret[hash] == nil {
			ret[hash] = []string{path}
		} else {
			ret[hash] = append(ret[hash], path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		// For the purposes of this program, leave it alone if we can't stat it
		return true
	}
	return fi.Mode().IsDir()
}

func hashFile(path string) (string, error) {
	hasher := sha256.New()
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func delQuery(files []string) {
	for i, file := range files {
		fmt.Println(i, file)
	}
	for {
		fmt.Printf("\nKeep which one? [0-%d, or blank for all] ", len(files)-1)
		scanner.Scan()
		res := scanner.Text()
		if res == "" {
			return
		}
		index, err := strconv.Atoi(res)
		if err == nil {
			if index < 0 || index >= len(files) {
				fmt.Println("Invalid index.")
			} else {
				for i, file := range files {
					if index == i {
						fmt.Printf("\t[+] %s\n", file)
					} else {
						fmt.Printf("\t[-] %s\n", file)
						delFile(file)
					}
				}
				return
			}
		}
	}
}

func delFile(file string) {
	os.Remove(file)
}
