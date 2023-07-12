package shortener

import "crypto/sha256"

func getSha256(s string) []byte {
	algo := sha256.New()
	algo.Write([]byte(s))

	return algo.Sum(nil)
}
