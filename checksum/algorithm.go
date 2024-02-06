package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"strings"
)

var ErrInvalidHashAlgorithm = errors.New("invalid hash algorithm")

var validHashAlgorithms = map[string]hash.Hash{
	"md5":    md5.New(),
	"sha1":   sha1.New(),
	"sha256": sha256.New(),
	"sha512": sha512.New(),
}

func verifyHashAlgorithm(alg string) error {
	_, ok := validHashAlgorithms[alg]
	if !ok {
		return ErrInvalidHashAlgorithm
	}
	return nil
}

func NewHashAlgorithm(alg string) (hash.Hash, error) {
	err := verifyHashAlgorithm(strings.ToLower(strings.TrimSpace(alg)))
	if err != nil {
		return nil, err
	}
	return validHashAlgorithms[alg], nil
}
