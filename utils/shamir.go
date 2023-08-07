package utils

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hashicorp/vault/shamir"
)

func splitSecret(secret string, parts, threshold int) ([]string, error) {

	var stringParts []string
	byteSecret := []byte(secret)

	byteParts, err := shamir.Split(byteSecret, parts, threshold)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to split secret: %v\n", err)
		return nil, err
	}

	for _, bytePart := range byteParts {
		str := bytes.NewBuffer(bytePart).String()
		stringParts = append(stringParts, str)
	}

	return stringParts, nil
}

func combineSecret(secrets []string) (string, error) {

	var byteParts [][]byte

	for _, stringPart := range secrets {
		bytePart := []byte(stringPart)
		byteParts = append(byteParts, bytePart)
	}

	secretByte, err := shamir.Combine(byteParts)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to combine secret: %v\n", err)
		return "", err
	}

	secret := bytes.NewBuffer(secretByte).String()

	return secret, nil
}
