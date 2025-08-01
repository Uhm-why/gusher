package hash

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func HashFile(fname string) (string, error) {

	file, err := os.Open(fname)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hasher := sha256.New()

	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil

}
