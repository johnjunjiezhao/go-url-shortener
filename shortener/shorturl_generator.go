package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(input []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(input)
	if err != nil {
		panic(err)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string, userID string) string {
	urlHashBytes := sha256Of(initialLink + userID)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
