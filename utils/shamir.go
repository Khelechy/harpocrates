package utils

import (
	"encoding/hex"

	"github.com/hashicorp/vault/shamir"
	"github.com/fatih/color"
)

func SplitSecret(secret string, parts, threshold int) ([]string, error) {

	var stringParts []string
	byteSecret := []byte(secret)

	byteParts, err := shamir.Split(byteSecret, parts, threshold)
	if err != nil {
		color.Red("Failed to split secret: %v\n", err)
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
		bytePart, _ := hex.DecodeString(stringPart)
		byteParts = append(byteParts, bytePart)
	}

	secretByte, err := shamir.Combine(byteParts)

	if err != nil {
		color.Red("Failed to combine secret: %v\n", err)
		return "", err
	}

	secret := string(secretByte)

	return secret, nil
}
