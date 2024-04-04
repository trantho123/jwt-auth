package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GetRandNum() (string, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(8999))
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}
