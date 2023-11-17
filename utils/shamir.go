package utils

import (
	"bytes"
	"encoding/hex"
	"log"

	"github.com/hashicorp/vault/shamir"
)

func SplitSecret(secret string, parts, threshold int) ([]string, error) {

	var stringParts []string
	byteSecret := []byte(secret)

	byteParts, err := shamir.Split(byteSecret, parts, threshold)
	if err != nil {
		log.Fatalf("Failed to split secret: %v\n", err)
		return nil, err
	}

	for _, bytePart := range byteParts {
		str := hex.EncodeToString(bytePart)
		stringParts = append(stringParts, str)
	}

	return stringParts, nil
}

func CombineSecret(secrets []string) (string, error) {

	var byteParts [][]byte

	for _, stringPart := range secrets {
		bytePart := []byte(stringPart)
		byteParts = append(byteParts, bytePart)
	}

	secretByte, err := shamir.Combine(byteParts)

	if err != nil {
		log.Fatalf("Failed to combine secret: %v\n", err)
		return "", err
	}

	secret := bytes.NewBuffer(secretByte).String()

	return secret, nil
}
