package auth

import (
	"crypto/rand"
	"math/big"
)

var (
	tokenCharacters   = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.-_")
	randomTokenLength = 16
	clientPrefix      = "C"
	randReader        = rand.Reader
)

func GenerateNotExistingToken(generateClientToken func() string, isTokenExists func(token string) bool) string {
	for {
		token := generateClientToken()
		if !isTokenExists(token) {
			return token
		}
	}
}

func GenerateClientToken() string {
	return generateRandomToken(clientPrefix)
}

func generateRandomToken(prefix string) string {
	return prefix + generateRandomString(randomTokenLength)
}

func randInt(n int) int {
	max := big.NewInt(int64(n))
	res, err := rand.Int(randReader, max)
	if err != nil {
		panic("Random resource is not available!!!")
	}
	return int(res.Int64())
}

func generateRandomString(length int) string {
	res := make([]byte, length)
	for i := range res {
		res[i] = tokenCharacters[randInt(len(tokenCharacters))]
	}
	return string(res)
}
