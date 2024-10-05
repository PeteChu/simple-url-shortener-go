package main

import (
	"crypto/sha256"
	"fmt"
)

func genUrlKey(url string) (result string) {
	result = fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
	return result[:5]
}
