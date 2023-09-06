package shortener

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/itchyny/base58-go"
)

func getSha256(s string) []byte {
	algo := sha256.New()
	algo.Write([]byte(s))

	return algo.Sum(nil)
}

func base58Encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	enc, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(enc)
}
