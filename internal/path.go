package internal

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// const charset = "abcdefghijklmnopqrstuvwxyz" +
// 	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const lowcharset = "abcdefghijklmnopqrstuvwxyz"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return RandomStringWithCharset(length, lowcharset)
}

func Exists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func MakeDirectoryIfNotExists(path string) error {

	if !Exists(path) {
		return os.MkdirAll(path, os.ModeDir|0755)
	}
	return nil
}

func DeleteDirectoryIfExists(path string) error {
	if Exists(path) {
		return os.RemoveAll(path)
	}
	return nil
}

func GetBinPath(tool string) string {
	path, err := exec.LookPath(tool)
	if err != nil {
		fmt.Printf("didn't find '%s' executable\n", tool)
		return ""
	}
	return path
}
